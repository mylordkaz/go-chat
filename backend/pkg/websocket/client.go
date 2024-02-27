package websocket

import (
	"encoding/json"
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
	Text string 	`json:"text"`
	Sender string 	`json:"sender"`

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
		// message := Message{Type: messageType, Body: string(p)}
		// c.Pool.Broadcast <- message
		// fmt.Println("Message Received: %+v\n", message)
		var receivedMessage Message
		err = json.Unmarshal(p, &receivedMessage)
		if err != nil {
			log.Println("Error decoding message:", err)
			continue
		}

		// Assuming 'Type' is set appropriately in the client's message
		receivedMessage.Type = messageType

		c.Pool.Broadcast <- receivedMessage
		fmt.Printf("Message Received: %+v\n", receivedMessage)
	}
}