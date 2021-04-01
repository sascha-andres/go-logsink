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
	"context"
	pb "github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

//lineSender emits the passed lines over gRPC
func lineSender(in <-chan string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("connect.address"), grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect: %s", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logrus.Fatalf("error closing: %s", err)
		}
	}()
	c := pb.NewLogTransferClient(conn)
	priority := int32(viper.GetInt("connect.priority"))
	for line := range in {
		_, err := c.SendLine(context.Background(), &pb.LineMessage{
			Line:     line,
			Priority: priority,
			Sequence: 0,
		})
		if err != nil {
			logrus.Warnf("problem sending line: %s", err)
		}
	}
}
