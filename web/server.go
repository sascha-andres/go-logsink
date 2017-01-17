package web

import (
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/sascha-andres/go-logsink/logsink"
)

// server is used to implement logsink.LogTransferServer.
type server struct {
	hub *hub
}

// SendLine implements logsink.SendLine
func (s *server) SendLine(ctx context.Context, in *pb.LineMessage) (*pb.LineResult, error) {
	s.hub.broadcast <- []byte(in.Line)
	return &pb.LineResult{Result: true}, nil
}

func (s *server) run() {
	s.hub = newHub()
	go s.hub.run()
	lis, err := net.Listen("tcp", viper.GetString("listen.bind"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterLogTransferServer(srv, s)
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
