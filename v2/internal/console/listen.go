package console

import (
	"github.com/sascha-andres/go-logsink/v2/internal/grpcserver"
	pb "github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/spf13/viper"
)

// Listen starts the Server
func Listen() {
	pipe := make(chan *pb.LineMessage)
	if viper.GetBool("listen.colored") {
		go printOutput(pipe, coloredPrinter)
	} else {
		go printOutput(pipe, nonColoredPrinter)
	}

	grpcserver.Listen(pipe)
}
