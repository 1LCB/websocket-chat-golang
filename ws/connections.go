package ws

import (
	"log/slog"
)

func HandleWsConnection(c *Client) {
	conn := c.conn
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()

		slog.Info("Message received")
		if err != nil {
			slog.Error(err.Error())
			break
		}
		
		BroadcastMessage(msg, c)
	}

	slog.Info("Connection of user " + c.username + " has been closed!")
}


func CloseAndRemoveConnection(client *Client, index int) {
	mu.Lock()
	defer mu.Unlock()

	client.conn.Close()

	clients = append(clients[:index], clients[index + 1:]...)

	slog.Info("Client removed successfully (" + client.username + ")")
}
