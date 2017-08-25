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

import "fmt"

type User struct {
	Id           int
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
	redmine_id   int
}

type UsersResponse struct {
	Users []User `json:"results"`
}

func (u *UsersResponse) Print() {
	for _, user := range u.Users {
		fmt.Printf("%s \"%s\" %s\n", user.FirstName, user.Username, user.LastName)
	}
}
