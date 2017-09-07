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
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var PerformancesColumns = []string{"day", "contract", "description", "duration", "multiplier", "type", "id"}
var PerformancesColumnsDefault = []bool{true, true, true, true, false, false, false}

type Performance struct {
	Id          int             `json:"id"`
	Type        PerformanceType `json:"type"`
	Timesheet   *Timesheet
	TimesheetId int `json:"timesheet"`
	ContractId  int `json:"contract"`
	Contract    *Contract
	Day         int    `json:"day"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	RateId      int    `json:"performance_type"`
	Rate        *PerformanceRate
}

func GetDefaultPerformancesColumns() string {
	var columns []string
	for i, defValue := range PerformancesColumnsDefault {
		if defValue {
			columns = append(columns, PerformancesColumns[i])
		}
	}
	return strings.Join(columns, ",")
}

type PerformancesList struct {
	Performances []Performance `json:"results"`
}

func (ps *PerformancesList) Fetch(c Client, t Timesheet) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_performances?timesheet=%d&page_size=9999", c.Endpoint, t.Id))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(ps)
	if err != nil {
		return err
	}

	contracts := &ContractsList{}
	err = contracts.Fetch(c)
	if err != nil {
		return err
	}

	rates := &PerformanceRatesList{}
	err = rates.Fetch(c)
	if err != nil {
		return err
	}

	for i := range ps.Performances {
		p := &ps.Performances[i]
		p.Timesheet = &t
		p.Contract = contracts.GetById(p.ContractId)
		p.Rate = rates.GetById(p.RateId)
		if err != nil {
			return err
		}
	}

	sort.SliceStable(ps.Performances, func(i, j int) bool {
		if ps.Performances[i].Day == ps.Performances[j].Day {
			return ps.Performances[i].Contract.PrettyLabel() < ps.Performances[j].Contract.PrettyLabel()
		} else {
			return ps.Performances[i].Day < ps.Performances[j].Day
		}
	})
	return nil
}

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

func (p *Performance) GetById(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_performances/%s/%d/", c.Endpoint, p.Type.String(), p.Id))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Performance) DeleteById(c Client) error {
	_, err := c.DeleteRequest(fmt.Sprintf("%s/v1/my_performances/%s/%d/", c.Endpoint, p.Type.String(), p.Id))
	if err != nil {
		return err
	}
	return nil
}

func (p *Performance) FetchContract(c Client) error {
	contract := &Contract{Id: p.ContractId}
	err := contract.GetById(c)
	if err != nil {
		return err
	}
	p.Contract = contract
	return nil
}

func (p *Performance) FetchRate(c Client) error {
	rate := &PerformanceRate{Id: p.RateId}
	err := rate.Fetch(c)
	if err != nil {
		return err
	}
	p.Rate = rate
	return nil
}

func (p *Performance) FetchTimesheet(c Client) error {
	timesheet := &Timesheet{Id: p.TimesheetId}
	err := timesheet.GetById(c)
	if err != nil {
		return err
	}
	p.Timesheet = timesheet
	return nil
}

func (ps *PerformancesList) PrettyPrint() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(60)
	table.SetHeader([]string{"Day", "Contract", "Description", "H", "Rate", "Type"})
	for _, p := range ps.Performances {
		table.Append([]string{strconv.Itoa(p.Day),
			p.Contract.Label, p.Description, p.Duration,
			p.Rate.Multiplier,
			p.Type.String(),
		})
	}
	table.Render()
}

func (p *Performance) PrettyPrint() {
	fmt.Printf("%d %s\n", p.Day, p.Description)
}

func (ps *PerformancesList) Porcelain() {
	for _, p := range ps.Performances {
		p.Porcelain()
	}
}

func (p *Performance) Sporcelain() string {
	return fmt.Sprintf("%d,%02d/%02d/%d,%s,%s,%s,%s,%s",
		p.Id, p.Day, p.Timesheet.Month, p.Timesheet.Year, p.Description,
		p.Contract.Label, p.Duration, p.Rate.Multiplier, p.Type.String())
}
func (p *Performance) Porcelain() {
	fmt.Println(p.Sporcelain())
}

func (ps *PerformancesList) PrettyPrintWithColumns(columns []string) {
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
func (p *PerformancesList) GetColumn(name string) string {
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

func (p *Performance) GetColumn(name string) string {
	switch name {
	case "id":
		return strconv.Itoa(p.Id)
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
