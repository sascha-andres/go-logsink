package client

import "fmt"

//lineFormatter formats a line before sending to server
func lineFormatter(in chan<- string) chan<- string {
	out := make(chan string)
	go func() {
		linePrefix := setupPrefix()
		for line := range in {
			if "" == linePrefix {
				out <- line
			} else {
				out <- fmt.Sprintf("[%s] %s", linePrefix, line)
			}
		}
		close(out)
	}()
	return out
}
