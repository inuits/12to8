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

// listRatesCmd represents the list rates command
var listRatesCmd = &cobra.Command{
	Use:   "rates",
	Short: "List all the rates",
	Run: func(cmd *cobra.Command, args []string) {
		pr := &api.PerformanceRatesList{}
		c := NewAPIClient()
		err := pr.Fetch(c)
		if err != nil {
			log.Fatal(err)
		}
		pr.PrettyPrint()
	},
}

func init() {
	listCmd.AddCommand(listRatesCmd)
}
