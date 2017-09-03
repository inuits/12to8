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
	"github.com/spf13/viper"
)

var project string
var multiplier string
var perf_type string

// newPerformanceCmd represents the newPerformance command
var newPerformanceCmd = &cobra.Command{
	Use:   "performance -p options dd[/MM[/YYYY]] H [DESCRIPTION]",
	Short: "Insert a new performance",
	Args:  validPerfAddArgs,
	Long: `Appends hours on the given day and project.

examples:
12to8 new performance 1 8

Contract can be specified on the CLI, via the config file ~/12to8.yml,
or via the env variable TWELVE_TO_EIGHT_CONTRACT.
It must follow the same syntax as in "12to8 list contracts".`,
	Run: func(cmd *cobra.Command, args []string) {
		contractLabel := viper.GetString("contract")
		if contractLabel == "" {
			log.Fatal("You must specify a contract in CLI, env variable or config file.")
		}
		perfTypeString := viper.GetString("type")
		if perfTypeString == "" {
			log.Fatal("You must specify a type in CLI, env variable or config file.")
		}

		c := NewApiClient()

		contracts := &api.ContractsList{}
		err := contracts.Fetch(c)
		contract, err := contracts.GetByLabel(contractLabel)
		if err != nil {
			log.Fatal(err)
		}
		if contract == nil {
			log.Fatalf("Contract %s not found", contractLabel)
		}

		rates := &api.PerformanceRatesList{}
		err = rates.Fetch(c)
		rate, err := rates.GetByMultiplier(multiplier)
		if err != nil {
			log.Fatal(err)
		}
		if rate == nil {
			log.Fatalf("Rate %s not found. Possible rates:\n%s", multiplier, rates.PrettyList())
		}

		var perfType api.PerformanceType
		switch perfTypeString {
		case "Activity":
			perfType = api.ActivityPerformance
		case "Standby":
			perfType = api.StandbyPerformance
		}

		day, month, year, err := getDaysMonthYearFromArg(args[0])
		if err != nil {
			log.Fatal(err)
		}

		timesheet := &api.Timesheet{
			Month: month,
			Year:  year,
		}

		err = timesheet.Get(c)
		if err != nil {
			log.Fatal(err)
		}

		var desc string
		if len(args) > 2 {
			desc = args[2]
		}

		performance := &api.Performance{
			Timesheet:   timesheet.Id,
			Day:         day,
			ContractId:  contract.Id,
			Description: desc,
			Type:        perfType,
			Duration:    args[1],
			RateId:      rate.Id,
		}

		err = performance.New(c)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	newCmd.AddCommand(newPerformanceCmd)

	newPerformanceCmd.Flags().StringVarP(&project, "contract", "c", "", "Contract to use")
	viper.BindPFlag("contract", newPerformanceCmd.Flags().Lookup("contract"))
	newPerformanceCmd.Flags().StringVarP(&perf_type, "type", "t", "Activity", "Type: Activity/Standby")
	viper.BindPFlag("type", newPerformanceCmd.Flags().Lookup("type"))
	newPerformanceCmd.Flags().StringVarP(&multiplier, "multiplier", "m", "1.00", "Multiplier")
}
