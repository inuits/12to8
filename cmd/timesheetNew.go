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
	"strconv"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

// timesheetNewCmd represents the timesheetNew command
var timesheetNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new timesheet",
	Long: `Create a new timesheet for the given month and year.
Example:

12to8 timesheet new 2017 09
12to8 timesheet new 2017 10 "My Timesheet"`,
	Args: newTimesheetArgs,
	Run: func(cmd *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		month, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		timesheet := &api.Timesheet{
			Month:  month,
			Year:   year,
			Status: "PENDING",
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

func init() {
	timesheetCmd.AddCommand(timesheetNewCmd)
}

func newTimesheetArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("requires exactly 2 args")
	}
	_, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("year must be a number")
	}
	month, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("month must be a number")
	}
	if month < 1 || month > 12 {
		return errors.New("invalid month range")
	}
	return nil
}
