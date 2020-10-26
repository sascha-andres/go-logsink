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
	web2 "github.com/sascha-andres/go-logsink/v2/internal/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start a server instance with a web interface",
	Long: `Use web to start a web server. Navigate with your favorite
browser to localhost:8080 ( change the binding definition )
to see the logs in your browser.

  go-logsink web --serve ":80" --bind ":50051"

If debug mode is enable you call open /debug/statsviz/ in your browser`,
	Run: func(cmd *cobra.Command, args []string) {
		handleLock(web2.Start)
	},
}

func init() {
	RootCmd.AddCommand(webCmd)
	webCmd.Flags().StringP("bind", "b", ":50051", "Provide bind definition")
	webCmd.Flags().StringP("serve", "s", ":8080", "Provide bind definition for web ui")
	webCmd.Flags().StringP("from-directory", "d", "", "provide a directory containing html files")
	viper.BindPFlag("web.bind", webCmd.Flags().Lookup("bind"))
	viper.BindPFlag("web.serve", webCmd.Flags().Lookup("serve"))
	viper.BindPFlag("web.directory", webCmd.Flags().Lookup("from-directory"))
}
