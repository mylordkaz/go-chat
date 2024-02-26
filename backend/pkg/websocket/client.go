package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)


type Client struct {
	ID string 					// Unique ID for particular connection
	Conn *websocket.Conn		// pointer to a websocket.Conn object
	Pool *Pool					// pointer to the pool which this client is part of
}

type Message struct {
	Type int   		`json:"type"`
	Body string 	`json:"body"`
}

// read method constently listen for new message coming through that connection
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Println("Message Received: %+V\n", message)
	}
}