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
var force bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "12to8",
	Short: "A client for 925r",
	Long:  "12to8 is a commandline application for the 925r, a free and open source time and leave tracking application.",
	BashCompletionFunction: bashCompletionFunc,
}

type logWriter struct {
}

// a writer that logs without dates
func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
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
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.12to8.yaml)")
	RootCmd.PersistentFlags().String("cache", "~/.cache/12to8", "config file (default is $HOME/.cache/12to8)")
	viper.BindPFlag("cache", RootCmd.PersistentFlags().Lookup("cache"))
	RootCmd.PersistentFlags().Bool("no-cache", false, "do not use cache, fetch from 925r as needed")
	viper.BindPFlag("no-cache", RootCmd.PersistentFlags().Lookup("no-cache"))
	RootCmd.PersistentFlags().StringP("user", "u", "", "username")
	viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user"))
	RootCmd.PersistentFlags().StringP("password", "p", "", "password")
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
	RootCmd.PersistentFlags().StringP("endpoint", "e", "", "API endpoint (without /v1)")
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

// NewAPIClient creates a new API client and populate its cache
// It gets endpoint, user, password from viper
func NewAPIClient() api.Client {
	username := viper.GetString("user")
	password := viper.GetString("password")
	endpoint := viper.GetString("endpoint")
	noCache := viper.GetBool("no-cache")
	cache := viper.GetString("cache")
	if endpoint == "" {
		log.Fatal("Endpoint is not set!")
	}
	c := api.Client{
		Username: username,
		Password: password,
		Endpoint: endpoint,
		NoCache:  noCache,
		CacheDir: cache,
	}
	err := c.FetchCache()
	if err != nil {
		log.Fatal(err)
	}
	return c
}
