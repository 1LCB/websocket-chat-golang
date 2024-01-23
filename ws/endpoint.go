package ws

import (
	"encoding/json"
	"log/slog"
	"net/http"

	gorilla "github.com/gorilla/websocket"
)

var upgrader = gorilla.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func send_response(w http.ResponseWriter, payload map[string]any, status_code int) {
	response, err := json.Marshal(payload)

	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(status_code)
	w.Write(response)
}

func WebsocketEndpoint(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	username := query.Get("username")

	if username == "" {
		send_response(w,
			map[string]any{
				"error": "you need to set a username first",
			}, http.StatusUnprocessableEntity,
		)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		send_response(w,
			map[string]any{
				"error": err.Error(),
			}, http.StatusBadRequest,
		)
		slog.Error(err.Error())
		return
	}

	slog.Info("New connection received!")

	client := NewClient(username, conn)
	go HandleWsConnection(client)
}
