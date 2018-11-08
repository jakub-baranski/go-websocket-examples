package main

import (
	"encoding/base64"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CatHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := GetConnection(w, r, nil)
	if err != nil {
		return
	}

	response, err := http.Get("https://cataas.com/cat")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	i, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	encoded := base64.StdEncoding.EncodeToString(i)

	for range time.Tick(5 * time.Second) {
		conn.WriteMessage(websocket.TextMessage, []byte(encoded))
	}

}
