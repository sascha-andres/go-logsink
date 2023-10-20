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

package grpcserver

import (
	pb "github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// Listen starts the Server
func Listen(out chan *pb.LineMessage) {
	logrus.Printf("Binding definition provided: %s\n", viper.GetString("listen.bind"))

	lis, err := net.Listen("tcp", viper.GetString("listen.bind"))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogTransferServer(s, NewServer(out))
	// Register reflection service on gRPC Server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
