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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Fetch the list of users",
	Long: `This command fetches the list of users
from 925r and displays it in a nice way.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		url := fmt.Sprintf("%s/v1/users?page_size=9999", endpoint)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Can prepare request: %v", err)
			os.Exit(1)
		}
		req.SetBasicAuth(username, password)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Can not fetch users list: %v", err)
			os.Exit(1)
		}
		if resp.StatusCode != 200 {
			log.Fatalf("Received non-200 status code while fetching %s: %d", url, resp.StatusCode)
			os.Exit(1)
		}
		users := api.UsersList{}
		json.NewDecoder(resp.Body).Decode(&users)
		users.PrettyPrint()
	},
}

func init() {
	RootCmd.AddCommand(usersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
