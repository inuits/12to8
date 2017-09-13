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

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

// timesheetCmd represents the timesheet command
var listCmd = &cobra.Command{
	Use:       "list [models]",
	Short:     "lists timesheets, performances, leaves...",
	ValidArgs: api.Models.List(),
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			return err
		}
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		m := api.Models.GetBySlug(args[0])
		c := NewAPIClient()
		c.FetchIfNeeded(m)
		m.PrettyPrint()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
