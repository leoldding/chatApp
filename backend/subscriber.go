package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func roomWS(w http.ResponseWriter, r *http.Request) {
	log.Println("ATTEMPTING TO CONNECT TO: " + r.URL.String())

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection to WebSocket: %f", err)
	}

	log.Print("SUCCESSFUL CONNECTION TO: " + r.URL.String())

	roomId := strings.Split(r.URL.String(), "/")[3]

	conn := &connection{ws: ws, sendMessage: make(chan []byte, 256)}
	sub := subscriber{conn: conn, roomId: roomId}
	pub.register <- sub
}

func mainWS(w http.ResponseWriter, r *http.Request) {
	log.Print("ATTEMPTING TO CONNECT TO: " + r.URL.String())

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection to WebSocket: %f", err)
	}

	log.Print("SUCCESSFUL CONNECTION TO: " + r.URL.String())
}
