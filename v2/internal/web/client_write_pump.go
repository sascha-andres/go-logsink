package web

import (
	"github.com/gorilla/websocket"
	"time"
)

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
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
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Enqueue(message)
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
