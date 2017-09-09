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
	"strconv"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// delete_performanceCmd represents the delete_performance command
var deletePerformanceCmd = &cobra.Command{
	Use:   "performance ID",
	Short: "delete a performance",
	Long: `Delete a performance by its id.
Use --force to avoid confirmation.`,
	Args: validIDArg,
	Run: func(cmd *cobra.Command, args []string) {
		perfTypeString := viper.GetString("type")
		if perfTypeString == "" {
			log.Fatal("You must specify a type in CLI, env variable or config file.")
		}
		var perfType api.PerformanceType
		switch perfTypeString {
		case "Activity":
			perfType = api.ActivityPerformance
		case "Standby":
			perfType = api.StandbyPerformance
		}

		c := NewApiClient()

		id, _ := strconv.Atoi(args[0])
		performance := &api.Performance{
			ID:   id,
			Type: perfType,
		}
		if !force {
			err := performance.GetByID(c)
			if err != nil {
				log.Fatal(err)
			}
			err = performance.FetchTimesheet(c)
			if err != nil {
				log.Fatal(err)
			}
			err = performance.FetchRate(c)
			if err != nil {
				log.Fatal(err)
			}
			err = performance.FetchContract(c)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Performance: %s.\nAre you sure you want to delete that performance? [y/N] ",
				performance.Sporcelain())
			var response string
			_, err = fmt.Scanln(&response)
			if err != nil {
				log.Fatal("Aborted by user")
			}
			if response != "yes" && response != "y" {
				log.Fatal("Aborted by user")
			}
		}

		err := performance.DeleteByID(c)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deletePerformanceCmd)

	deletePerformanceCmd.Flags().BoolVarP(&force, "force", "f", false, "Do not prompt before deleting")

	// type
	deletePerformanceCmd.Flags().StringVarP(&perfType, "type", "t", "Activity", "Type: Activity/Standby")
	// autocomplete
	annotation := make(map[string][]string)
	annotation[cobra.BashCompCustom] = []string{"__12to8_comp_activity"}
	c := deletePerformanceCmd.Flags().Lookup("type")
	c.Annotations = annotation
	viper.BindPFlag("type", c)
}
