// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-logsink",
	Short: "go-logsink is a simplistic log aggregator",
	Long: `You can use go-logsink to combine multiple log streams
into one. For example you can combine multiple tails into one
output.

To do this start a server and connect any number of clients.`,
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

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-logsink.yaml)")
	RootCmd.PersistentFlags().StringP("lockfile", "", "", "Specify a lockfile")
	RootCmd.PersistentFlags().BoolP("debug", "", false, "Enable debug mode (statsviz)")

	viper.BindPFlag("lockfile", RootCmd.PersistentFlags().Lookup("lockfile"))
	viper.BindPFlag("debug", RootCmd.PersistentFlags().Lookup("debug"))

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-logsink") // name of config file (without extension)
	viper.AddConfigPath("$HOME")       // adding home directory as first search path
	viper.AddConfigPath(".")           // look in currect directory for config
	viper.AutomaticEnv()               // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if !viper.GetBool("connect.pass-through") {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
