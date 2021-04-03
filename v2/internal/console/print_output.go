package console

import "fmt"

//printOutput prints the received line to console
func printOutput(in <-chan string) {
	for line := range in {
		fmt.Println(line)
	}
}
