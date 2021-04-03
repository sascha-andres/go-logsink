package console

import (
	"github.com/sascha-andres/go-logsink/v2/internal/grpcserver"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Listen starts the Server
func Listen() {
	logrus.Printf("Binding definition provided: %s\n", viper.GetString("listen.bind"))

	pipe := make(chan string)
	go printOutput(pipe)

	grpcserver.Listen(pipe)
}
