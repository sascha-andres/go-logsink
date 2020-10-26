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
	client2 "github.com/sascha-andres/go-logsink/v2/internal/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a go-logsink server and forward stdin",
	Long: `This command is used to connect to a go-logsink server.
Call it to forward data piped ito this application to the server.

If you want to filter the function maust be named filter and return a bool:

    def filter(line):
      return line.startswith("a")

The above filter function will not print lines starting with the letter a (lowercase)`,
	Run: func(cmd *cobra.Command, args []string) {
		handleLock(client2.Connect)
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringP("address", "a", "localhost:50051", "Provide server address")
	connectCmd.Flags().StringP("prefix", "p", "", "Provide a prefix for each line")
	connectCmd.Flags().IntP("priority", "", 0, "Priority of message")
	connectCmd.Flags().BoolP("pass-through", "", false, "Print lines to stdout")
	connectCmd.Flags().StringP("filter-function", "", "", "Provide path to starlark file to filter lines")

	_ = viper.BindPFlag("connect.address", connectCmd.Flags().Lookup("address"))
	_ = viper.BindPFlag("connect.prefix", connectCmd.Flags().Lookup("prefix"))
	_ = viper.BindPFlag("connect.priority", connectCmd.Flags().Lookup("priority"))
	_ = viper.BindPFlag("connect.pass-through", connectCmd.Flags().Lookup("pass-through"))
	_ = viper.BindPFlag("connect.filter-function", connectCmd.Flags().Lookup("filter-function"))
}
