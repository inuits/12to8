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
	bash_completion_func = `
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

__12to8_contracts_comp(){
	local IFS=$'\n'
	COMPREPLY=( $( compgen -W "$(12to8 list contracts 2>/dev/null)" -- "$cur" ) )
	local i=0
	local r
	for r in ${COMPREPLY[@]}
	do
		if [[ "${cur:0:1}" == "'" ]]; then
			COMPREPLY[$i]="${r//\'/\'}"
		elif [[ "${cur:0:1}" == '"' ]]; then
			COMPREPLY[$i]="${r//\"/\\\"}"
		else
			COMPREPLY[$i]="\"${r//\"/\\\"}\""
		fi
		let i++
	done
	return 0
}

__custom_func() {
    case ${last_command} in
        12to8_new_timesheet)
            __12to8_new_timesheet_comp
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
	Use:       "completion SHELL",
	Short:     "Output shell completion code for the specified shell.",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"bash"},
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
	c := NewApiClient()
	err := contracts.Fetch(c)
	if err != nil {
		log.Fatal(err)
	}
	contracts.PrettyPrint()
}
func init() {
	RootCmd.AddCommand(autocompleteCmd)
}
