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
	"log"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

// list_performancesCmd represents the list_performances command
var listPerformancesCmd = &cobra.Command{
	Use:   "performances [MM[/YYYY]]",
	Short: "List the performances for the given timesheet",
	Args:  validTimesheetArgs,
	Long: `This command will show you all your performances for a
given timesheet.

If no timesheet is given, take the current month.`,
	Run: func(cmd *cobra.Command, args []string) {
		timesheet := fetchTimesheetFromArgs(args)
		performances := &api.PerformancesList{}
		c := NewApiClient()
		err := performances.Fetch(c, *timesheet)
		if err != nil {
			log.Fatal(err)
		}
		performances.PrettyPrint()
	},
}

func init() {
	listCmd.AddCommand(listPerformancesCmd)
}
