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
	"fmt"
	"log"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

// timesheetReleaseCmd represents the timesheetRelease command
var releaseTimesheetCmd = &cobra.Command{
	Use:   "timesheet [MM[/YYYY]]",
	Short: "Releate a timesheet to pending state",
	Long: `Release a timesheet for approval. Once you have done this,
you can no longer make changes to the timesheet.`,
	Args: validTimesheetArgs,
	Run: func(cmd *cobra.Command, args []string) {
		monthSpec := ""
		if len(args) == 1 {
			monthSpec = args[0]
		}
		month, year, err := getMonthYearFromArg(monthSpec)
		if err != nil {
			log.Fatal(err)
		}
		timesheet := &api.Timesheet{
			Month: month,
			Year:  year,
		}
		c := NewApiClient()
		err = timesheet.Get(c)
		if err != nil {
			log.Fatal(err)
		}
		if !force {
			fmt.Printf("Are you sure you want to release the timesheet %s?\n",
				timesheet.Name())
			fmt.Printf("No changes are possible after that. [y/N] ")
			var response string
			_, err = fmt.Scanln(&response)
			if err != nil {
				log.Fatal("Aborted by user")
			}
			if response != "yes" && response != "y" {
				log.Fatal("Aborted by user")
			}
		}
		err = timesheet.Release(c)
		if err != nil {
			log.Fatal(err)
		}
		timesheet.PrettyPrint()
	},
}

func init() {
	releaseCmd.AddCommand(releaseTimesheetCmd)

	releaseTimesheetCmd.Flags().BoolVarP(&force, "force", "f", false, "Do not prompt before releasing")
}
