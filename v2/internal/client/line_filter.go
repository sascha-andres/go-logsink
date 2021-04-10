// Copyright © 2021 Sascha Andres <sascha.andres@outlook.com>
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

package client

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//lineFilter reduces lines to those which are not filtered out
func lineFilter(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		err := setupFilter()
		if err != nil {
			logrus.Fatalf("could not setupPrefix filter: %s", err)
			close(out)
			return
		}
		passThrough := viper.GetBool("connect.pass-through")

		for line := range in {
			if filtered(line) {
				continue
			}
			out <- line
			if passThrough {
				fmt.Println(line)
			}
		}
		close(out)
	}()

	return out
}

