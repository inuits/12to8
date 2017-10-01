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

// LeaveTypes stores the performances rates we get from the cache or the remote server
var LeaveTypes = &LeaveTypesList{}

// LeaveType represents a Performance Rate in the ninetofiver api.
type LeaveType struct {
	ID    int    `json:"id"`
	Label string `json:"label"`
}

// LeaveTypesList represents a list of Performances Rates as we get them from the server.
type LeaveTypesList struct {
	LeaveTypes []LeaveType `json:"results"`
}

func (lt *LeaveTypesList) apiURL() string {
	return "v1/performance_types"
}

// Slug is used to represent the model in cli
func (lt *LeaveTypesList) Slug() string {
	return "leave-types"
}

func (lt *LeaveTypesList) augment(c *Client) error {
	return nil
}

func (lt *LeaveTypesList) isEmpty() bool {
	return len(lt.LeaveTypes) == 0
}

// GetByID fills a performance given its ID.
func (lt *LeaveTypesList) GetByID(id int) *LeaveType {
	for i := range lt.LeaveTypes {
		p := lt.LeaveTypes[i]
		if p.ID == id {
			return &p
		}
	}
	return nil
}

// PrettyPrint prints the performances rates in a nice way.
func (lt *LeaveTypesList) PrettyPrint(columns []string) {
	for _, p := range lt.LeaveTypes {
		p.PrettyPrint()
	}
}

// PrettyPrint prints the performance rate in a nice way.
func (p *LeaveType) PrettyPrint() {
	fmt.Printf("%s\n", p.Label)
}

func (lt *LeaveTypesList) extraFetchParameters(c Client, args []string) string {
	return ""
}

// HasPorcelain returns false because performances rates do not support PorcelainPrettyPrint
func (lt *LeaveTypesList) HasPorcelain() bool {
	return false
}

// GetColumns returns an empty list because performances rates are not displayed as a table
func (lt *LeaveTypesList) GetColumns() []string {
	return *new([]string)
}

// GetDefaultColumns returns an empty list because Performance rates are not displayed as a table
func (lt *LeaveTypesList) GetDefaultColumns() []string {
	return *new([]string)
}

// PorcelainPrettyPrint does nothing for this model
func (lt *LeaveTypesList) PorcelainPrettyPrint() {
	return
}

func init() {
	cache.register(LeaveTypes)
	Models.register(LeaveTypes)
}
