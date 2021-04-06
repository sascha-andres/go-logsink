package web

import (
	"encoding/json"
	"fmt"
	"github.com/sascha-andres/go-logsink/v2/logsink"
	"github.com/sirupsen/logrus"
)

//processOutput prints the received line to console
func (s *server) processOutput(in <-chan *logsink.LineMessage) {
	var cnt int64

	for line := range in {
		cnt++
		numberOfLines.Inc()
		obj := lineType{
			Line:     line.Line,
			Priority: line.Priority, // TODO Priority
			Key:      fmt.Sprintf("%d", cnt),
		}
		if line.Line == "" {
			obj.Line = " "
		}
		data, err := json.Marshal(obj)
		if err != nil {
			logrus.Errorf("error creating message for websocket: %s", err)
		} else {
			s.hub.broadcast <- data
		}
	}

}
