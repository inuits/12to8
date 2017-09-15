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
	"strings"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

var columns string
var porcelain bool

// timesheetCmd represents the timesheet command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list timesheets, performances, leaves...",
}

func init() {
	RootCmd.AddCommand(listCmd)
	for i, m := range api.Models.Models {
		mCmd := &cobra.Command{
			Use:   fmt.Sprintf("%s", api.Models.Models[i].Slug()),
			Short: fmt.Sprintf("list the %s", api.Models.Models[i].Slug()),
			Args: func(cmd *cobra.Command, args []string) error {
				var invalidColumns []string
				for _, column := range strings.Split(columns, ",") {
					if column == "" {
						continue
					}
					found := false
					for _, validColumn := range api.PerformancesColumns {
						if validColumn == column {
							found = true
						}
					}
					if !found {
						invalidColumns = append(invalidColumns, column)
					}
				}
				if len(invalidColumns) > 0 {
					return fmt.Errorf("invalid columns: %s", strings.Join(invalidColumns, ", "))
				}
				return nil
			},
			Run: listRun(m),
		}
		listCmd.AddCommand(mCmd)
		if m.HasPorcelain() {
			mCmd.Flags().BoolVarP(&porcelain, "porcelain", "P", false, "porcelain (usable in scripts) output")
		}
		if len(m.GetColumns()) > 0 {
			mCmd.Flags().StringVarP(&columns, "columns", "C", "", "comma-separated columns to be displayed")
		}
	}

}

func listRun(m api.ModelList) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		c := NewAPIClient()
		c.FetchIfNeeded(m, args)
		if porcelain {
			m.PorcelainPrettyPrint()
			return
		}
		if columns != "" {
			m.PrettyPrint(strings.Split(columns, ","))
			return
		}
		m.PrettyPrint(m.GetDefaultColumns())
	}
}
