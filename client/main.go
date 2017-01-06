package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	pb "github.com/sascha-andres/go-logsink/logsink"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLogTransferClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res, err := c.SendLine(context.Background(), &pb.LineMessage{Line: scanner.Text()})
		if !res.Result {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
