package ws

import (
	"log/slog"
	"sync"

	gorilla "github.com/gorilla/websocket"
)

type Client struct {
	username string
	conn     *gorilla.Conn
}

var mu = sync.Mutex{}
var clients = []*Client{}

func NewClient(username string, conn *gorilla.Conn) *Client {
	c := &Client{
		username: username,
		conn:     conn,
	}

	clients = append(clients, c)

	return c
}

func BroadcastMessage(msg []byte, client_sender *Client) { //vai ser uma go routine
	mu.Lock()
	defer mu.Unlock()

	for i, client := range clients {
		err := client.conn.WriteMessage(gorilla.TextMessage, msg)
		if err != nil {
			CloseAndRemoveConnection(client, i)
		}

		slog.Info("Message sent successfully to " + client.username)
	}
}