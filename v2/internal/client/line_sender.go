package client

import (
	"github.com/sascha-andres/go-logsink/v2/logsink"
	pb "github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//lineSender emits the passed lines over gRPC
func lineSender(in chan<- string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("connect.address"), grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect: %s", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}()
	c := logsink.NewLogTransferClient(conn)
	client, err := c.SendLine(context.Background())
	if err != nil {
		logrus.Panicf("error creating client to send log entries: %s", err)
	}
	for line := range in {
		err = client.Send(&pb.LineMessage{Line: line, Priority: int32(viper.GetInt("connect.priority"))})
		if nil != err {
			log.Fatal(err)
		}
	}
	res, err := client.CloseAndRecv()
	if !(nil != res && res.Result) || nil != err {
		logrus.Fatal(err)
	}
}
