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
	"fmt"
	"log"
	"strings"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

var columns string
var porcelain bool

// timesheetCmd represents the timesheet command
var listCmd = &cobra.Command{
	Use:       "list MODEL [args...]",
	Short:     "lists timesheets, performances, leaves...",
	ValidArgs: api.Models.List(),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		found := false
		for _, model := range api.Models.List() {
			if model == args[0] {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("Not a model: %s", args[0])
		}
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
	Run: func(cmd *cobra.Command, args []string) {
		m := api.Models.GetBySlug(args[0])
		c := NewAPIClient()
		var model string
		if len(args) > 0 {
			model = args[0]
			args = args[1:]
		}
		c.FetchIfNeeded(m, args)
		if porcelain {
			if !m.HasPorcelain() {
				log.Fatalf("%s do not support porcelain", model)
			}
			m.PorcelainPrettyPrint()
			return
		}
		if columns != "" {
			if len(m.GetColumns()) == 0 {
				log.Fatalf("%s do not support columns", model)
			}
			cols := strings.Split(columns, ",")
			m.PrettyPrint(cols)
			return
		}
		m.PrettyPrint(m.GetDefaultColumns())
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&columns, "columns", "C", "", "comma-separated columns to be displayed")
	listCmd.Flags().BoolVarP(&porcelain, "porcelain", "P", false, "porcelain (usable in scripts) output")
}
