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
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
)

//lineProducer reads lines and emits them to the pipeline
func lineProducer(reader io.Reader) <-chan string {
	out := make(chan string)

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Warnf("reading standard input: %s", err)
		}
		close(out)
	}()

	return out
}

