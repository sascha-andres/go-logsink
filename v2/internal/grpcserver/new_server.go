package grpcserver

import pb "github.com/sascha-andres/go-logsink/v2/logsink"

///NewServer constructs a Server
func NewServer(out chan<- *pb.LineMessage) *Server {
	return &Server{
		output: out,
		UnimplementedLogTransferServer: pb.UnimplementedLogTransferServer{},
	}
}
