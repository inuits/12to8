// Copyright © 2017 Julien Pivotto <roidelapluie@inuits.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/inuits/12to8/helpers"
	"github.com/olekukonko/tablewriter"
)

// Performances stores the performances as we get them from the ninetofiver server
var Performances = &PerformancesList{}

// PerformancesColumns are all the columns we can show for performances.
var PerformancesColumns = []string{"day", "contract", "description", "duration", "multiplier", "type", "id"}

// PerformancesColumnsDefault is a list of booleans. true means column will be shown by default.
// Each boolean represents the column at the same place in the PerformancesColumns list.
var PerformancesColumnsDefault = []bool{true, true, true, true, false, false, false}

// Performance represents one performance we get in the ninetofoiver app
type Performance struct {
	ID          int             `json:"id"`
	Type        PerformanceType `json:"type"`
	Timesheet   *Timesheet
	TimesheetID int `json:"timesheet"`
	ContractID  int `json:"contract"`
	Contract    *Contract
	Day         int    `json:"day"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	RateID      int    `json:"performance_type"`
	Rate        *PerformanceRate
}

// HasPorcelain returns true because contracts do support PorcelainPrettyPrint
func (ps *PerformancesList) HasPorcelain() bool {
	return true
}

// GetColumns returns all the available columns in the performances tabular
func (ps *PerformancesList) GetColumns() []string {
	return PerformancesColumns
}

// GetDefaultColumns returns the list of default columns for performances
func (ps *PerformancesList) GetDefaultColumns() []string {
	var columns []string
	for i, defValue := range PerformancesColumnsDefault {
		if defValue {
			columns = append(columns, PerformancesColumns[i])
		}
	}
	return columns
}

// PerformancesList represents a list of performances we get from the ninetofiver server
type PerformancesList struct {
	Performances []Performance `json:"results"`
}

func (ps *PerformancesList) apiURL() string {
	return "v1/my_performances"
}

func (p *Performance) apiURL() string {
	return "v1/my_performances"
}

func (ps *PerformancesList) augment(c *Client) error {
	for i := range ps.Performances {
		p := &ps.Performances[i]
		c.FetchIfNeeded(timesheets, *new([]string))
		p.Timesheet = timesheets.GetByID(p.TimesheetID)
		p.Contract = Contracts.GetByID(p.ContractID)
		p.Rate = PerformancesRates.GetByID(p.RateID)
		// if err != nil {
		// return err
		// }
	}

	sort.SliceStable(ps.Performances, func(i, j int) bool {
		if ps.Performances[i].Day == ps.Performances[j].Day {
			return ps.Performances[i].Contract.PrettyLabel() < ps.Performances[j].Contract.PrettyLabel()
		}
		return ps.Performances[i].Day < ps.Performances[j].Day
	})

	return nil

}

func (ps *PerformancesList) isEmpty() bool {
	return len(ps.Performances) == 0
}

// Slug is used to represent the model in cli
func (ps *PerformancesList) Slug() string {
	return "performances"
}

// Slug is used to represent the model in cli
func (p *Performance) Slug() string {
	return "performance"
}

// SetID sets the id of the performance
func (p *Performance) SetID(i int) {
	p.ID = i
}

// GetID returns the id of the performance
func (p *Performance) GetID() int {
	return p.ID
}

// DeleteArg returns the args needed to delete a performance
func (p *Performance) DeleteArg() string {
	return fmt.Sprintf("%s/%d", p.Type.String(), p.ID)
}

// Augment populates extra fields of a performance
func (p *Performance) Augment(c *Client) {
	c.FetchIfNeeded(timesheets, *new([]string))
	p.Timesheet = timesheets.GetByID(p.TimesheetID)
	p.Contract = Contracts.GetByID(p.ContractID)
	p.Rate = PerformancesRates.GetByID(p.RateID)
}

// New creates a new performance on the server
func (p *Performance) New(c Client) error {
	resp, err := c.PostRequest(fmt.Sprintf("%s/v1/my_performances/%s/", c.Endpoint, p.Type.String()), p)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return err
	}
	return nil
}

// GetByID gets the performance from the server given its it
func (p *Performance) GetByID(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_performances/%s/%d/", c.Endpoint, p.Type.String(), p.ID))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return err
	}
	return nil
}

// FetchContract fills the contract field of this performance
func (p *Performance) FetchContract(c Client) error {
	p.Contract = Contracts.GetByID(p.ContractID)
	return nil
}

// FetchRate fills the rate field of this performance
func (p *Performance) FetchRate(c Client) error {
	p.Rate = PerformancesRates.GetByID(p.RateID)
	return nil
}

// FetchTimesheet fills the timesheet field of this performance
func (p *Performance) FetchTimesheet(c Client) error {
	timesheet := &Timesheet{ID: p.TimesheetID}
	err := timesheet.GetByID(c)
	if err != nil {
		return err
	}
	p.Timesheet = timesheet
	return nil
}

// PorcelainPrettyPrint prints out the porcelain version of the performances
func (ps *PerformancesList) PorcelainPrettyPrint() {
	for _, p := range ps.Performances {
		p.Porcelain()
	}
}

// Sporcelain creates a parsable string for the performance.
// Useful for scripting.
func (p *Performance) Sporcelain() string {
	return fmt.Sprintf("%d,%02d/%02d/%d,%s,%s,%s,%s,%s",
		p.ID, p.Day, p.Timesheet.Month, p.Timesheet.Year, p.Description,
		p.Contract.Label, p.Duration, p.Rate.Multiplier, p.Type.String())
}

// Porcelain prints out the porcelain version of the performance
func (p *Performance) Porcelain() {
	fmt.Println(p.Sporcelain())
}

// PrettyPrint prints a list of performances as a table
// which columns specified as parameter
func (ps *PerformancesList) PrettyPrint(columns []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(60)
	var header []string
	for _, c := range columns {
		header = append(header, ps.GetColumn(c))
	}
	table.SetHeader(header)
	for _, p := range ps.Performances {
		var row []string
		for _, c := range columns {
			row = append(row, p.GetColumn(c))
		}
		table.Append(row)
	}
	table.Render()
}

// GetColumn returns the header of a column for performances
func (ps *PerformancesList) GetColumn(name string) string {
	switch name {
	case "id":
		return "ID"
	case "day":
		return "Day"
	case "contract":
		return "Contract"
	case "description":
		return "Description"
	case "duration":
		return "H"
	case "multiplier":
		return "x"
	case "type":
		return "Kind"
	}
	return ""
}

// PrettyPrint prints the performance
func (p *Performance) PrettyPrint() {
	fmt.Printf("%d - %02d/%02d/%d %sh %s [%s]\n",
		p.ID, p.Day, p.Timesheet.Month, p.Timesheet.Year, p.Duration, p.Description,
		p.Contract.Label)
}

// GetColumn returns the content of a column for a performance
func (p *Performance) GetColumn(name string) string {
	switch name {
	case "id":
		return strconv.Itoa(p.ID)
	case "day":
		return strconv.Itoa(p.Day)
	case "contract":
		return p.Contract.Label
	case "description":
		return p.Description
	case "duration":
		return p.Duration
	case "multiplier":
		return p.Rate.Multiplier
	case "type":
		return p.Type.String()
	}
	return ""
}

func (ps *PerformancesList) extraFetchParameters(c Client, args []string) string {
	monthSpec := ""
	if len(args) == 1 {
		monthSpec = args[0]
	}
	month, year, err := helpers.GetMonthYearFromArg(monthSpec)
	if err != nil {
		log.Fatal(err)
	}
	timesheet := &Timesheet{
		Month: month,
		Year:  year,
	}
	err = timesheet.Get(c)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("&timesheet=%d", timesheet.ID)
}

func init() {
	Models.register(Performances, &Performance{})
}
