package console

import pb "github.com/sascha-andres/go-logsink/v2/logsink"

//printOutput prints the received line to console
func printOutput(in <-chan *pb.LineMessage, printer LinePrinter) {
	for input := range in {
		printer(input)
	}
}
