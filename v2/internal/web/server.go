package web

import (
	"github.com/sascha-andres/go-logsink/v2/internal/grpcserver"
	pb "github.com/sascha-andres/go-logsink/v2/logsink"
)

// server is used to implement logsink.LogTransferServer.
type (
	server struct {
		pb.UnimplementedLogTransferServer
		hub           *hub
		numberOfLines int64
	}
	lineType struct {
		Line     string
		Priority int32
		Key      string
	}
)

func (s *server) run(out chan *pb.LineMessage) {
	go s.hub.run()
	grpcserver.Listen(out)
}
