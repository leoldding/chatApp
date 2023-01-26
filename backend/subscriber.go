package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func connWS(w http.ResponseWriter, r *http.Request, roomId string) {
	log.Print("Connected to room: " + roomId)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection to WebSocket: %f", err)
	}

	conn := &connection{ws: ws, sendMessage: make(chan []byte, 256)}
	sub := subscriber{conn: conn, roomId: roomId}
	pub.register <- sub
}
