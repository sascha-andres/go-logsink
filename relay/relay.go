package relay

import (
	"log"
	"net"

	pb "github.com/sascha-andres/go-logsink/logsink"
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
func Relay(port, address string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c = pb.NewLogTransferClient(conn)

	lis, err := net.Listen("tcp", port)
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
