package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
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
	go sub.readMsgs()
	go sub.writeMsgs()
}

func mainWS(w http.ResponseWriter, r *http.Request) {
	log.Print("ATTEMPTING TO CONNECT TO: " + r.URL.String())

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection to WebSocket: %f", err)
	}

	log.Print("SUCCESSFUL CONNECTION TO: " + r.URL.String())

	conn := &connection{ws: ws, sendMessage: make(chan []byte, 256)}

	ticker := time.NewTicker(30 * time.Second)

	for range ticker.C {
		if err := conn.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			return
		}
	}
}

func (sub subscriber) readMsgs() {
	c := sub.conn

	defer func() {
		pub.deregister <- sub
		c.ws.Close()
	}()

	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			log.Printf("Websocket error: %f", err)
			break
		}

		m := message{msg, sub.roomId}
		pub.broadcast <- m
	}
}

func (sub *subscriber) writeMsgs() {
	c := sub.conn

	defer func() {
		c.ws.Close()
	}()

	ticker := time.NewTicker(30 * time.Second)

	for {
		select {
		case msg, ok := <-c.sendMessage:
			if !ok {
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.ws.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
