package client

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
)

//lineProducer reads lines and emits them to the pipeline
func lineProducer(reader io.Reader) chan<- string {
	out := make(chan string)

	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Warnf("reading standard input: %s", err)
		}
		close(out)
	}()

	return out
}

