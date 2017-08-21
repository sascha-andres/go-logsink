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

package relay

import (
	"net"

	log "github.com/sirupsen/logrus"

	pb "github.com/sascha-andres/go-logsink/logsink"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	c pb.LogTransferClient
)

// server is used to implement logsink.LogTransferServer.
type relayServer struct{}

// SendLine implements logsink.SendLine
func (s *relayServer) SendLine(ctx context.Context, in *pb.LineMessage) (*pb.LineResult, error) {
	c.SendLine(context.Background(), &pb.LineMessage{Line: in.Line})
	return &pb.LineResult{Result: true}, nil
}

// Relay starts a server and connects to a client
func Relay() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("relay.address"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = pb.NewLogTransferClient(conn)

	lis, err := net.Listen("tcp", viper.GetString("relay.bind"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogTransferServer(s, &relayServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
