package server

import (
	"log"
	"net"

	"fmt"

	pb "github.com/sascha-andres/go-logsink/logsink"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement logsink.LogTransferServer.
type server struct{}

// SendLine implements logsink.SendLine
func (s *server) SendLine(ctx context.Context, in *pb.LineMessage) (*pb.LineResult, error) {
	fmt.Println(in.Line)
	return &pb.LineResult{Result: true}, nil
}

// Listen starts the server
func Listen() {
	fmt.Printf("Binding definition provided: %s\n", viper.GetString("listen.bind"))

	lis, err := net.Listen("tcp", viper.GetString("listen.bind"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogTransferServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
