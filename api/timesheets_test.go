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

func ExampleTimesheet_PrettyPrint() {
	t := &Timesheet{
		Year:   2016,
		Month:  7,
		Status: "PENDING",
	}
	t.PrettyPrint()
	// Output: July 2016 [PENDING]
}

func ExampleTimesheets_PrettyPrint() {
	t := &TimesheetsList{
		Timesheets: []Timesheet{
			Timesheet{
				Year:   2007,
				Month:  8,
				Status: "APPROVED",
			},
			Timesheet{
				Year:   2016,
				Month:  7,
				Status: "PENDING",
			},
		},
	}
	t.PrettyPrint()
	// Output:
	// August 2007 [APPROVED]
	// July 2016 [PENDING]
}
