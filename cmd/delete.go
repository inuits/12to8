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
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete timesheets, performances, leaves...",
}

func init() {
	RootCmd.AddCommand(deleteCmd)
	for i, m := range api.Models.IndividualModels {
		mCmd := &cobra.Command{
			Use:   fmt.Sprintf("%s ID", api.Models.IndividualModels[i].Slug()),
			Short: fmt.Sprintf("delete the %s", api.Models.IndividualModels[i].Slug()),
			Run:   deleteRun(m),
		}
		mCmd.Flags().BoolVarP(&force, "force", "f", false, "Do not prompt before deleting")
		deleteCmd.AddCommand(mCmd)
	}
}

func deleteRun(m api.Model) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		c := NewAPIClient()

		id, _ := strconv.Atoi(args[0])
		m.SetID(id)
		err := c.GetByID(m)
		if err != nil {
			log.Fatal(err)
		}
		if !force {
			m.Augment(&c)
			m.PrettyPrint()
			fmt.Printf("Are you sure you want to delete that %s? [y/N] ",
				m.Slug())
			var response string
			_, err := fmt.Scanln(&response)
			if err != nil {
				log.Fatal("Aborted by user")
			}
			if response != "yes" && response != "y" {
				log.Fatal("Aborted by user")
			}
		}

		err = c.DeleteByID(m)
		if err != nil {
			log.Fatal(err)
		}
	}
}
