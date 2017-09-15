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
	"fmt"
)

// Contracts stores the contracts we have in cache or fetched from the server.
var Contracts = &ContractsList{}

// Contract is a contract as seen in the ninetofiver api.
type Contract struct {
	ID         int    `json:"id"`
	Label      string `json:"label"`
	CustomerID int    `json:"customer"`
	Customer   *Company
}

// ContractsList is a list of contrats as seen in the ninetofiver api.
type ContractsList struct {
	Contracts []Contract `json:"results"`
}

func (cs *ContractsList) apiURL() string {
	return "v1/my_contracts"
}

func (cs *ContractsList) slug() string {
	return "contracts"
}

func (cs *ContractsList) isEmpty() bool {
	return len(cs.Contracts) == 0
}

// GetByID returns the contract with the given id
func (cs *ContractsList) GetByID(id int) *Contract {
	for _, c := range cs.Contracts {
		if c.ID == id {
			return &c
		}
	}
	return nil
}

// GetByLabel returns the contract with the given label
func (cs *ContractsList) GetByLabel(label string) *Contract {
	for _, c := range cs.Contracts {
		if c.PrettyLabel() == label {
			return &c
		}
	}
	return nil
}

func (cs *ContractsList) augment(c *Client) error {
	for i := range cs.Contracts {
		co := &cs.Contracts[i]
		for _, customer := range Companies.Companies {
			if customer.ID == co.CustomerID {
				co.Customer = &customer
				break
			}
		}
	}
	return nil
}

// PrettyPrint prints contracts in a nice way to the console
func (cs *ContractsList) PrettyPrint(columns []string) {
	for _, c := range cs.Contracts {
		c.PrettyPrint()
	}
}

// PrettyLabel returns the label of a contract that can be used
// in CLI etc.. to identify a contract. It contains the customer.
func (c *Contract) PrettyLabel() string {
	if c.Customer == nil {
		return fmt.Sprintf("%s [%d]", c.Label, c.CustomerID)
	}
	return fmt.Sprintf("%s [%s]", c.Label, c.Customer.Name)
}

// PrettyPrint prints contract in a nice way to the console
func (c *Contract) PrettyPrint() {
	fmt.Println(c.PrettyLabel())
}

func (cs *ContractsList) extraFetchParameters(c Client, args []string) string {
	return ""
}

// HasPorcelain returns false because contracts do not support PorcelainPrettyPrint
func (cs *ContractsList) HasPorcelain() bool {
	return false
}

// GetColumns returns an empty list because contracts are not displayed as a table
func (cs *ContractsList) GetColumns() []string {
	return *new([]string)
}

// GetDefaultColumns returns an empty list because contracts are not displayed as a table
func (cs *ContractsList) GetDefaultColumns() []string {
	return *new([]string)
}

// PorcelainPrettyPrint does nothing for this model
func (cs *ContractsList) PorcelainPrettyPrint() {
	return
}

func init() {
	cache.register(Contracts)
	Models.register(Contracts)
}
