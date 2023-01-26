package main

import (
	"github.com/gorilla/websocket"
)

type message struct {
	text   []byte
	roomId string
}

type connection struct {
	ws          *websocket.Conn
	sendMessage chan []byte
}

type subscriber struct {
	conn   *connection
	roomId string
}

type publisher struct {
	broadcast  chan message
	register   chan subscriber
	deregister chan subscriber
	rooms      map[string]map[*connection]bool
}
