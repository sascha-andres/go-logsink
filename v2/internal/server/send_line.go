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

package server

import (
	"github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/sirupsen/logrus"
)

//SendLine implements logsink.SendLine
func (s *server) SendLine(stream logsink.LogTransfer_SendLineServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			logrus.Warnf("error reading request: %s", err)
			return stream.SendAndClose(&logsink.LineResult{
				Result: true,
			})
		}
		logrus.Println(in.Line)
	}
	logrus.Warnf("this code should not be reached")
	return nil
}
