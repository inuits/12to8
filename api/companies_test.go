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

func ExampleCompany_PrettyPrint() {
	t := &Company{
		Name:    "Inuits",
		Country: "BE",
	}
	t.PrettyPrint()
	// Output: Inuits [BE]
}

func ExampleCompaniesList_PrettyPrint() {
	t := &CompaniesList{
		Companies: []Company{
			Company{
				Name:    "Inuits",
				Country: "BE",
			},
			Company{
				Name:    "Eskimo",
				Country: "NL",
			},
		},
	}
	t.PrettyPrint()
	// Output:
	// Inuits [BE]
	// Eskimo [NL]
}
