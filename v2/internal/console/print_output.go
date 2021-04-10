package console

import "fmt"
import pb "github.com/sascha-andres/go-logsink/v2/logsink"

//printOutput prints the received line to console
func printOutput(in <-chan *pb.LineMessage) {
	for input := range in {
		fmt.Println(input.Line)
	}
}
