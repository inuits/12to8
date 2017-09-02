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
	"os"

	"github.com/inuits/12to8/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var username string
var password string
var endpoint string
var force bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "12to8",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	BashCompletionFunction: bash_completion_func,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.12to8.yaml)")
	RootCmd.PersistentFlags().StringVarP(&username, "user", "u", "", "username")
	viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user"))
	RootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password")
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
	RootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "", "API endpoint (without /v1)")
	viper.BindPFlag("endpoint", RootCmd.PersistentFlags().Lookup("endpoint"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".12to8")          // name of config file (without extension)
	viper.AddConfigPath(os.Getenv("HOME")) // adding home directory as first search path
	viper.SetEnvPrefix("twelve_to_eight")  // env variables can't start with a number
	viper.AutomaticEnv()                   // read in environment variables that match
	viper.ReadInConfig()
}

func NewApiClient() api.Client {
	username := viper.GetString("user")
	password := viper.GetString("password")
	endpoint := viper.GetString("endpoint")
	if endpoint == "" {
		log.Fatal("Endpoint is not set!")
	}
	return api.Client{
		Username: username,
		Password: password,
		Endpoint: endpoint,
	}
}
