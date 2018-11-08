package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)


// TODO: keep and return recent chat history for channel

const ChannelName = "channel1"

func ChatHandler(w http.ResponseWriter, r *http.Request) {

	// Get WebSocket connection
	conn, err := GetConnection(w, r, nil)
	if err != nil {
		return
	}

	// Create redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Subscribe to channel
	pubsub := client.Subscribe(ChannelName)
	ch := pubsub.Channel()
	// Test subscription
	_, err = pubsub.Receive()
	if err != nil {
		panic(err)
	}
	wmc := make(chan *webSocketMessage)
	defer close(wmc)
	go messageReader(conn, wmc)

	for {

		select {
		case pub := <-ch:
			conn.WriteMessage(websocket.TextMessage, []byte(pub.Payload))
		case received := <-wmc:
			if received.err != nil {
				panic(err)
			}
			now := time.Now()
			newMsg := fmt.Sprintf("%d-%d-%d: %s", now.Hour(), now.Minute(), now.Second(), received.msg)
			client.Append(ChannelName, newMsg)
			val := client.Get(ChannelName)
			fmt.Println(val)
			client.Publish(ChannelName, string(newMsg))
		}
	}

}

type webSocketMessage struct {
	msgType int
	msg []byte
	err error
}

func messageReader(conn *websocket.Conn, ch chan *webSocketMessage ) {
	for {
		t, msg, err := conn.ReadMessage()
		m := &webSocketMessage{t, msg, err}
		if err != nil {
			break
		}
		ch <- m
	}
}



