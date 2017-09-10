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

var Contracts = &ContractsList{}

type Contract struct {
	ID         int    `json:"id"`
	Label      string `json:"label"`
	CustomerID int    `json:"customer"`
	Customer   *Company
}

type ContractsList struct {
	Contracts []Contract `json:"results"`
}

func (cs *ContractsList) apiURL() string {
	return "v1/my_contracts"
}

func (cs *ContractsList) slug() string {
	return "contracts"
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

// GetByID returns the contract with the given id
func (cs *ContractsList) GetByLabel(label string) *Contract {
	for _, c := range cs.Contracts {
		if c.PrettyLabel() == label {
			return &c
		}
	}
	return nil
}

func (cs *ContractsList) augment() error {
	for i := range cs.Contracts {
		co := &cs.Contracts[i]
		for _, customer := range companies.Companies {
			if customer.ID == co.CustomerID {
				co.Customer = &customer
				break
			}
		}
	}
	return nil
}

func (cs *ContractsList) PrettyPrint() {
	for _, c := range cs.Contracts {
		c.PrettyPrint()
	}
}

func (c *Contract) PrettyLabel() string {
	if c.Customer == nil {
		return fmt.Sprintf("%s [%d]", c.Label, c.CustomerID)
	}
	return fmt.Sprintf("%s [%s]", c.Label, c.Customer.Name)
}

func (c *Contract) PrettyPrint() {
	fmt.Println(c.PrettyLabel())
}

func init() {
	cache.register(Contracts)
}
