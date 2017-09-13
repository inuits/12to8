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
	"fmt"
)

// PerformancesRates stores the performances rates we get from the cache or the remote server
var PerformancesRates = &PerformanceRatesList{}

// PerformanceRate represents a Performance Rate in the ninetofiver api.
type PerformanceRate struct {
	ID         int    `json:"id"`
	Label      string `json:"label"`
	Multiplier string `json:"multiplier"`
}

// PerformanceRatesList represents a list of Performances Rates as we get them from the server.
type PerformanceRatesList struct {
	PerformanceRates []PerformanceRate `json:"results"`
}

func (pr *PerformanceRatesList) apiURL() string {
	return "v1/performance_types"
}

func (pr *PerformanceRatesList) slug() string {
	return "rates"
}

func (pr *PerformanceRatesList) augment(c *Client) error {
	return nil
}

func (pr *PerformanceRatesList) isEmpty() bool {
	return len(pr.PerformanceRates) == 0
}

// GetByMultiplier fills a performance given its multiplier.
// It will error if multiple performances rates are found with the same multiplier.
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

// GetByID fills a performance given its ID.
func (pr *PerformanceRatesList) GetByID(id int) *PerformanceRate {
	for i := range pr.PerformanceRates {
		p := pr.PerformanceRates[i]
		if p.ID == id {
			return &p
		}
	}
	return nil
}

// PrettyList returns a short list of performances rates, used in error messages
// (when a user specifies an unknown multiplier)
func (p *PerformanceRate) PrettyList() string {
	return fmt.Sprintf(" * %s : %s\n", p.Multiplier, p.Label)
}

// PrettyList returns a string with all the performances rates displayed
// as a short list.
func (pr *PerformanceRatesList) PrettyList() string {
	var buffer bytes.Buffer
	for _, p := range pr.PerformanceRates {
		buffer.WriteString(p.PrettyList())
	}
	return buffer.String()
}

// PrettyPrint prints the performances rates in a nice way.
func (pr *PerformanceRatesList) PrettyPrint(columns []string) {
	for _, p := range pr.PerformanceRates {
		p.PrettyPrint()
	}
}

// ShortPrint prints all the multipliers of the performance rates in the list, and only that.
func (pr *PerformanceRatesList) ShortPrint() {
	for _, p := range pr.PerformanceRates {
		p.ShortPrint()
	}
}

// ShortPrint the multiplier of the performance rate, and only that.
func (p *PerformanceRate) ShortPrint() {
	fmt.Println(p.Multiplier)
}

// PrettyPrint prints the performance rate in a nice way.
func (p *PerformanceRate) PrettyPrint() {
	fmt.Printf("%s [%s]\n", p.Multiplier, p.Label)
}

func (pr *PerformanceRatesList) extraFetchParameters(c Client, args []string) string {
	return ""
}

// HasPorcelain returns false because performances rates do not support PorcelainPrettyPrint
func (pr *PerformanceRatesList) HasPorcelain() bool {
	return false
}

// GetColumns returns an empty list because performances rates are not displayed as a table
func (pr *PerformanceRatesList) GetColumns() []string {
	return *new([]string)
}

// GetDefaultColumns returns an empty list because Performance rates are not displayed as a table
func (pr *PerformanceRatesList) GetDefaultColumns() []string {
	return *new([]string)
}

// PorcelainPrettyPrint does nothing for this model
func (pr *PerformanceRatesList) PorcelainPrettyPrint() {
	return
}

func init() {
	cache.register(PerformancesRates)
	Models.register(PerformancesRates)
}
