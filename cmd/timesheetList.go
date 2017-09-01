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
	"os"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

// timesheetListCmd represents the timesheetList command
var timesheetListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		timesheets := &api.TimesheetsList{}
		c := NewApiClient()
		err := timesheets.Fetch(c)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		timesheets.PrettyPrint()
	},
}

func init() {
	timesheetCmd.AddCommand(timesheetListCmd)
}
