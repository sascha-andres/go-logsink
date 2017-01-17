package client

import (
	"bufio"
	"fmt"
	"log"
	"os"

	pb "github.com/sascha-andres/go-logsink/logsink"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	linePrefix string
)

// Connect is used to connect to a go-logsink server
func Connect() {

	fmt.Printf("Connecting to %s\n", viper.GetString("connect.address"))
	linePrefix = viper.GetString("connect.prefix")
	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("connect.address"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLogTransferClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var (
			res *pb.LineResult
			err error
		)
		if "" == linePrefix {
			res, err = c.SendLine(context.Background(), &pb.LineMessage{Line: scanner.Text()})
		} else {
			res, err = c.SendLine(context.Background(), &pb.LineMessage{Line: fmt.Sprintf("[%s] %s", linePrefix, scanner.Text())})
		}
		if !res.Result || nil != err {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
