package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ping", ping)

	http.HandleFunc("/ws", mainWS)

	http.HandleFunc("/ws/", roomWS)

	go pub.publish()
	http.ListenAndServe(":8080", nil)
	return
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
