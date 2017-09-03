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
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Performance struct {
	Id          int             `json:"id"`
	Type        PerformanceType `json:"type"`
	Timesheet   int             `json:"timesheet"`
	ContractId  int             `json:"contract"`
	Contract    *Contract
	Day         int    `json:"day"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	RateId      int    `json:"performance_type"`
	Rate        *PerformanceRate
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
		p.Contract = contracts.GetById(p.ContractId)
		p.Rate = rates.GetById(p.RateId)
		if err != nil {
			return err
		}
	}
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

func (p *Performance) FetchContract(c Client) error {
	contract := &Contract{Id: p.ContractId}
	err := contract.GetById(c)
	if err != nil {
		return err
	}
	p.Contract = contract
	return nil
}

func (ps *PerformancesList) PrettyPrint() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColWidth(60)
	table.SetHeader([]string{"Day", "Project", "Description", "H", "Rate", "Type"})
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
