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
	"io"
	"os"
)

//getReader constructs a reader from stdin or file
func getReader() (io.Reader, error) {
	var (
		reader io.Reader
		err error
	)
	var fileName = viper.GetString("connect.file")
	if "" != fileName {
		reader, err = os.Open(fileName)
		if err != nil {
			logrus.Fatalf("error opening file: %s", err)
			return nil, err
		}
	} else {
		reader = os.Stdin
	}
	return reader, err
}
