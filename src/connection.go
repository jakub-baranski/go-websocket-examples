package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func GetConnection(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// Trust all origins
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	return upgrader.Upgrade(w, r, nil)
}