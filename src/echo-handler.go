package main

import (
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := GetConnection(w, r, nil)
	if err != nil {
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, append(msg, []byte(", you say? - resended with love from server")... ) )
	}

}
