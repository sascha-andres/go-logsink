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
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//setupPrefix returns prefix to use
func setupPrefix() string {
	if !viper.GetBool("connect.pass-through") {
		logrus.Printf("Connecting to %s\n", viper.GetString("connect.address"))
	}
	return viper.GetString("connect.prefix")
}
