package console

import (
	"github.com/sascha-andres/go-logsink/v2/internal/grpcserver"
	pb "github.com/sascha-andres/go-logsink/v2/logsink"
)

// Listen starts the Server
func Listen() {
	pipe := make(chan *pb.LineMessage)
	go printOutput(pipe)

	grpcserver.Listen(pipe)
}
