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

package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

// timesheetNewCmd represents the timesheetNew command
var newTimesheetCmd = &cobra.Command{
	Use:   "timesheet [MM[/YYYY]]",
	Short: "Create a new timesheet",
	Long: `Create a new timesheet for the given month and year.
Example:

12to8 timesheet new
12to8 timesheet new 10
12to8 timesheet new 9/2017
`,
	Args: newTimesheetArgs,
	Run: func(cmd *cobra.Command, args []string) {
		monthSpec := ""
		if len(args) == 1 {
			monthSpec = args[0]
		}
		month, year, err := getMonthYearFromArg(monthSpec)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		timesheet := &api.Timesheet{
			Month:  month,
			Year:   year,
			Status: "ACTIVE",
		}
		c := NewApiClient()
		err = timesheet.New(c)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		timesheet.PrettyPrint()
	},
}

func newTimesheetArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return errors.New("takes at most one argument")
	}
	if len(args) == 1 {
		_, _, err := getMonthYearFromArg(args[0])
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	newCmd.AddCommand(newTimesheetCmd)
}
