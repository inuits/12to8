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

// Users is used to cache the list of Users
var Users = &UsersList{}

// User represents a ninetofiver user
type User struct {
	ID           int
	Username     string
	Email        string
	Groups       []string
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	DisplayLabel string
	IsActive     bool
	Country      string
	Gender       string
	BirthDate    string
	JoinDate     string
	RedmineID    int `json:"redmine_id"`
}

// UsersList is a list of ninetofiver users
type UsersList struct {
	Users []User `json:"results"`
}

// Slug is used to represent the model in cli
func (users *UsersList) Slug() string {
	return "users"
}

func (users *UsersList) apiURL() string {
	return "v1/users"
}

func (users *UsersList) augment(c *Client) error {
	return nil
}

func (users *UsersList) isEmpty() bool {
	return len(users.Users) == 0
}

func (users *UsersList) extraFetchParameters(c Client, args []string) string {
	return ""
}

// PrettyPrint prints users in a nice way to the console
func (users *UsersList) PrettyPrint(columns []string) {
	for _, u := range users.Users {
		u.PrettyPrint()
	}
}

// HasPorcelain returns false because users do not support PorcelainPrettyPrint
func (users *UsersList) HasPorcelain() bool {
	return false
}

// GetColumns returns an empty list because users are not displayed as a table
func (users *UsersList) GetColumns() []string {
	return *new([]string)
}

// GetDefaultColumns returns an empty list because users are not displayed as a table
func (users *UsersList) GetDefaultColumns() []string {
	return *new([]string)
}

// PorcelainPrettyPrint does nothing for this model
func (users *UsersList) PorcelainPrettyPrint() {
	return
}

// PrettyPrint prints user in a nice way to the console
func (u *User) PrettyPrint() {
	fmt.Printf("%s \"%s\" %s\n", u.FirstName, u.Username, u.LastName)
}

func init() {
	Models.register(Users)
}
