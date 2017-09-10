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
)

var companies = &CompaniesList{}

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type CompaniesList struct {
	Companies []Company `json:"results"`
}

func (cs *CompaniesList) apiURL() string {
	return "v1/companies"
}

func (cs *CompaniesList) slug() string {
	return "companies"
}

func (cs *CompaniesList) augment() error {
	return nil
}

// Get returns the Company from the server
func (c *Company) Get(client Client) error {
	cs := &CompaniesList{}
	resp, err := client.GetRequest(fmt.Sprintf("%s/v1/companies/?name=%s", client.Endpoint, c.Name))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(cs)
	if err != nil {
		return err
	}
	if len(cs.Companies) != 1 {
		return fmt.Errorf("Expected 1 company, got %d", len(cs.Companies))
	}
	*c = cs.Companies[0]
	return nil
}

// GetByID returns the Company from the server
func (c *Company) GetByID(client Client) error {
	resp, err := client.GetRequest(fmt.Sprintf("%s/v1/companies/%d/", client.Endpoint, c.ID))

	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(c)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CompaniesList) PrettyPrint() {
	for _, c := range cs.Companies {
		c.PrettyPrint()
	}
}

func (c *Company) PrettyPrint() {
	fmt.Printf("%s [%s]\n", c.Name, c.Country)
}

func init() {
	cache.register(companies)
}
