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
	"bytes"
	"fmt"
	"log"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

const (
	bashCompletionFunc = `
__12to8_new_timesheet_comp(){
    if [[ ${#nouns[@]} -eq 0 ]]; then
        COMPREPLY=( $( compgen -W "$(date -d "$(date -d "$(date +%Y-%m-01) - 1 day")" +%m/%Y) $(date +%m/%Y) $(date -d "$(date +"%Y-%m-01") 31 days" +%m/%Y)" -- "$cur" ) )
        return 0
    fi
    if [[ ${#nouns[@]} -ge 1 ]]; then
			COMPREPLY=(  $(compgen -W "" -- "$cur" ) )
        return 0
	fi
}

__12to8_comp_activity(){
	COMPREPLY=( $( compgen -W "activity standby" -- "$cur" ) )
}

__12to8_comp_hours(){
	COMPREPLY=( $( compgen -W "8.0" -- "$cur" ) )
}
__12to8_comp_close_dates(){
	COMPREPLY=( $( compgen -W "$(for i in {-5..5}; do date -d "$i day" +%d/%m/%Y; done)" -- "$cur" ) )
}

__12to8_new_performance_comp(){
    if [[ ${#nouns[@]} -eq 0 ]]; then
		__12to8_comp_close_dates
        return 0
    fi
    if [[ ${#nouns[@]} -eq 1 ]]; then
		__12to8_comp_hours
        return 0
	fi
}

__12to8_comp(){
	local IFS=$'\n'
	COMPREPLY=( $( compgen -W "$(12to8 completion $1 2>/dev/null)" -- "$cur" ) )
	local i=0
	local r
	for r in ${COMPREPLY[@]}
	do
		if [[ "${cur:0:1}" == "'" ]]; then
			COMPREPLY[$i]="${r//\'/\'}"
		elif [[ "${cur:0:1}" == '"' ]]; then
			COMPREPLY[$i]="${r//\"/\\\"}"
		else
			case "$r" in
			*\ *)
			COMPREPLY[$i]="\"${r//\"/\\\"}\"" ;;
			*\&*)
			COMPREPLY[$i]="\"${r//\"/\\\"}\"" ;;
			*\(*)
			COMPREPLY[$i]="\"${r//\"/\\\"}\""  ;;
			*\)*)
			COMPREPLY[$i]="\"${r//\"/\\\"}\""  ;;
			esac
		fi
		let i++
	done
	return 0
}

__custom_func() {
    case ${last_command} in
        12to8_new_timesheet | 12to8_release_timesheet)
            __12to8_new_timesheet_comp
            return
            ;;
		12to8_new_performance)
			__12to8_new_performance_comp
			return
			;;
        *)
            ;;
    esac
}
`
)

// autocompleteCmd represents the autocomplete command
var autocompleteCmd = &cobra.Command{
	Use:        "completion SHELL",
	Short:      "Output shell completion code for the specified shell.",
	Args:       cobra.RangeArgs(1, 2),
	ValidArgs:  []string{"bash"},
	ArgAliases: []string{"contracts", "rates", "performance_types"},
	Long: `To enable bash completion, run the following command or
add it to your ~/.bashrc

. <(12to8 completion bash)

or compile it to a static file:

12to8 completion bash > ~/.12to8.complete
echo . ~/.12to8.complete >> ~/.bashrc
. ~/.bashrc
`,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			bashComplete()
		case "contracts":
			contractsComplete()
		case "rates":
			ratesComplete()
		default:
			log.Fatal("Unknown shell")
		}
	},
}

func bashComplete() {
	var out bytes.Buffer
	RootCmd.GenBashCompletion(&out)
	fmt.Print(out.String())
}

// listContractsCmd represents the list contracts command
func contractsComplete() {
	contracts := &api.ContractsList{}
	c := NewAPIClient()
	err := contracts.Fetch(c)
	if err != nil {
		log.Fatal(err)
	}
	contracts.PrettyPrint()
}
func init() {
	RootCmd.AddCommand(autocompleteCmd)
}

// listRatesCmd represents the list rates command
func ratesComplete() {
	pr := &api.PerformanceRatesList{}
	c := NewAPIClient()
	err := c.FetchList(pr)
	if err != nil {
		log.Fatal(err)
	}
	pr.ShortPrint()
}
