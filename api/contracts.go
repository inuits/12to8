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
	"errors"
	"fmt"
)

type Contract struct {
	Id         int    `json:"id"`
	Label      string `json:"label"`
	CustomerId int    `json:"customer"`
	Customer   *Company
}

type ContractsList struct {
	Contracts []Contract `json:"results"`
}

func (cs *ContractsList) Fetch(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_contracts?page_size=9999", c.Endpoint))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(cs)
	if err != nil {
		return err
	}
	for i := range cs.Contracts {
		co := &cs.Contracts[i]
		err = co.FetchCustomer(c)
		if err != nil {
			return err
		}
	}
	return nil
}

// Get returns the Contract from the server
func (co *Contract) Get(c Client) error {
	cs := &ContractsList{}
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_contracts/?label=%s", c.Endpoint, co.Label))
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(cs)
	if err != nil {
		return err
	}
	if len(cs.Contracts) != 1 {
		return errors.New(fmt.Sprintf("Expected 1 contract, got %d", len(cs.Contracts)))
	}
	*co = cs.Contracts[0]
	err = co.FetchCustomer(c)
	if err != nil {
		return err
	}
	return nil
}

func (co *Contract) FetchCustomer(c Client) error {
	customer := &Company{Id: co.CustomerId}
	err := customer.GetById(c)
	if err != nil {
		return err
	}
	co.Customer = customer
	return nil
}

func (cs *ContractsList) GetById(id int) *Contract {
	for i := range cs.Contracts {
		c := cs.Contracts[i]
		if c.Id == id {
			return &c
		}
	}
	return nil
}

func (cs *ContractsList) GetByLabel(label string) (*Contract, error) {
	var contract *Contract
	for i := range cs.Contracts {
		c := cs.Contracts[i]
		if c.PrettyLabel() == label {
			if contract != nil {
				return nil, errors.New(fmt.Sprintf("Found 2 contracts with label %s", label))
			}
			contract = &c
		}
	}
	return contract, nil
}

// GetById returns the Company from the server
func (co *Contract) GetById(c Client) error {
	resp, err := c.GetRequest(fmt.Sprintf("%s/v1/my_contracts/%d/", c.Endpoint, co.Id))

	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(co)
	if err != nil {
		return err
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
		return fmt.Sprintf("%s [%d]", c.Label, c.CustomerId)
	}
	return fmt.Sprintf("%s [%s]", c.Label, c.Customer.Name)
}

func (c *Contract) PrettyPrint() {
	fmt.Println(c.PrettyLabel())
}
