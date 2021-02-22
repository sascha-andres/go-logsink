package web

import (
	"github.com/gorilla/websocket"
	"time"
)

//publishQueueContentToWebsocket takes the content of the queue and publishes
//to the websocket
func (c *client) publishQueueContentToWebsocket() {
	go func() {
		for {
			if c.HasElements() {
				w, err := c.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					break
				}
				result, data := c.Dequeue()
				if !result {
					continue
				}
				w.Write(data)

				if err := w.Close(); err != nil {
					break
				}
				continue
			}
			time.Sleep(100 * time.Microsecond)
		}
	}()
}
