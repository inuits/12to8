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

	"github.com/spf13/cobra"
)

// autocompleteCmd represents the autocomplete command
var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "Generate Bash autocomplete",
	Long: `This command generates the bash completion.

To enable it, simply run:

	. <(12to8 autocomplete)
	
You can also add that to your ~/.bashrc file.`,
	Run: func(cmd *cobra.Command, args []string) {
		var out bytes.Buffer
		RootCmd.GenBashCompletion(&out)
		fmt.Print(out.String())
	},
}

func init() {
	RootCmd.AddCommand(autocompleteCmd)
}
