package web

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// serveWebSocket handles websocket requests from the peer.
func serveWebSocket(hub *hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Println(err)
		return
	}
	client := &client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	go client.writePump()
	client.readPump()
}
