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
	"bytes"
	"encoding/json"
	"fmt"
)

type PerformanceRate struct {
	Id         int    `json:"id"`
	Label      string `json:"label"`
	Multiplier string `json:"multiplier"`
}

type PerformanceRatesList struct {
	PerformanceRates []PerformanceRate `json:"results"`
}

func (pr *PerformanceRatesList) Fetch(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/performance_types/?page_size=9999", c.Endpoint))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(pr)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PerformanceRatesList) GetByMultiplier(multiplier string) (*PerformanceRate, error) {
	var rate *PerformanceRate
	for i := range pr.PerformanceRates {
		p := pr.PerformanceRates[i]
		if p.Multiplier == multiplier {
			if rate != nil {
				return nil, fmt.Errorf("Found 2 rates with multiplier %s", multiplier)
			}
			rate = &p
		}
	}
	return rate, nil
}

func (pr *PerformanceRatesList) GetById(id int) *PerformanceRate {
	for i := range pr.PerformanceRates {
		p := pr.PerformanceRates[i]
		if p.Id == id {
			return &p
		}
	}
	return nil
}

func (p *PerformanceRate) Fetch(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/performance_types/%d/", c.Endpoint, p.Id))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return err
	}
	return nil
}

func (p *PerformanceRate) PrettyList() string {
	return fmt.Sprintf(" * %s : %s\n", p.Multiplier, p.Label)
}

func (pr *PerformanceRatesList) PrettyList() string {
	var buffer bytes.Buffer
	for _, p := range pr.PerformanceRates {
		buffer.WriteString(p.PrettyList())
	}
	return buffer.String()
}

func (pr *PerformanceRatesList) PrettyPrint() {
	for _, p := range pr.PerformanceRates {
		p.PrettyPrint()
	}
}

func (pr *PerformanceRatesList) ShortPrint() {
	for _, p := range pr.PerformanceRates {
		p.ShortPrint()
	}
}

func (p *PerformanceRate) ShortPrint() {
	fmt.Println(p.Multiplier)
}

func (p *PerformanceRate) PrettyPrint() {
	fmt.Printf("%s [%s]\n", p.Multiplier, p.Label)
}
