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

import log "github.com/sirupsen/logrus"

//Connect is used to connect to a go-logsink server
func Connect() {
	reader, err := getReader()
	if err != nil {
		log.Fatalf("error opening reader: %s", err)
	}

	lines := lineProducer(reader)
	filtered := lineFilter(lines)
	formatted := lineFormatter(filtered)
	lineSender(formatted)
}

