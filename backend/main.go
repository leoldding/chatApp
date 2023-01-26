package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/ping", ping)

	http.HandleFunc("/room/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("HANDLE FUNCTION")
		roomId := strings.Split(r.URL.String(), "/")[2]
		connWS(w, r, roomId)
	})

	go pub.publish()
	http.ListenAndServe(":8080", nil)
	return
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
