package web

import (
	"github.com/gorilla/websocket"
	"sync"
)

//NewClient creates a new client info struct
func NewClient(hub *hub, conn *websocket.Conn) (*client, error) {
	return &client{
		mutex:     sync.RWMutex{},
		hub:       hub,
		conn:      conn,
		send:      make(chan []byte, 1024),
		queueData: make([][]byte, 0),
	}, nil
}
