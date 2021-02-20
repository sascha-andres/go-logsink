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

package server

import (
	"github.com/arl/statsviz"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	pb "github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement logsink.LogTransferServer.
type server struct {
	pb.UnimplementedLogTransferServer
}

// SendLine implements logsink.SendLine
func (s *server) SendLine(stream pb.LogTransfer_SendLineServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(in.Line)
		stream.Send(&pb.LineResult{
			Result:               true,
			Sequence:             in.Sequence,
		})
	}
}

// Listen starts the server
func Listen() {
	log.Printf("Binding definition provided: %s\n", viper.GetString("listen.bind"))

	if viper.GetBool("debug") {
		go startDebug()
	}

	lis, err := net.Listen("tcp", viper.GetString("listen.bind"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogTransferServer(s, &server{
		UnimplementedLogTransferServer: pb.UnimplementedLogTransferServer{},
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startDebug() {
	r := mux.NewRouter()
	r.Methods("GET").Path("/debug/statsviz/ws").Name("GET /debug/statsviz/ws").HandlerFunc(statsviz.Ws)
	r.Methods("GET").PathPrefix("/debug/statsviz/").Name("GET /debug/statsviz/").Handler(statsviz.Index)

	h := &http.Server{Addr: viper.GetString("web.serve"), Handler: handlers.CORS()(handlers.ProxyHeaders(r))}
	if err := h.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
