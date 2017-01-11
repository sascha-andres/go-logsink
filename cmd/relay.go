// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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

	"github.com/sascha-andres/go-logsink/relay"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var relayCmd = &cobra.Command{
	Use:   "relay",
	Short: "Start a server that forwards messages",
	Long: `Instead of dumping incoming messages a relay forwards
the messages to another go-logsink instance`,
	Run: func(cmd *cobra.Command, args []string) {
		address := viper.GetString("address")
		if "" == address {
			log.Fatalf("You have to provide the address flag")
		}
		fmt.Printf("Connecting to %s\n", address)
		bind := viper.GetString("bind")
		fmt.Printf("Binding definition provided: %s\n", bind)
		relay.Relay(bind, address)
	},
}

func init() {
	RootCmd.AddCommand(relayCmd)

	relayCmd.Flags().StringP("bind", "b", ":50051", "Binding definition")
	relayCmd.Flags().StringP("address", "a", "", "Address to connect to")
	viper.BindPFlag("bind", listenCmd.Flags().Lookup("bind"))
	viper.BindPFlag("address", listenCmd.Flags().Lookup("address"))
}
