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

func (users *UsersList) slug() string {
	return "users"
}

func (users *UsersList) apiURL() string {
	return "v1/users"
}

func (users *UsersList) augment() error {
	return nil
}

// PrettyPrint prints users in a nice way to the console
func (users *UsersList) PrettyPrint() {
	for _, u := range users.Users {
		u.PrettyPrint()
	}
}

// PrettyPrint prints user in a nice way to the console
func (u *User) PrettyPrint() {
	fmt.Printf("%s \"%s\" %s\n", u.FirstName, u.Username, u.LastName)
}
