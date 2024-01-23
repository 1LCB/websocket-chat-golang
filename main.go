package main

import (
	"golang-chat/ws"
	"log/slog"
	"net/http"
)

func main() {
	address := "0.0.0.0:80"
	http.HandleFunc("/ws", ws.WebsocketEndpoint)

	slog.Info("Listening on " + address)

	if err := http.ListenAndServe(address, nil); err != nil{
		panic(err)
	}
}