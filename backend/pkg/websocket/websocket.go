package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Read & Write bufferSize to define an Upgrader (from http endpoint to ws endpoint)
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// check origin of our connection (now: allow any connection)
	CheckOrigin: func(r *http.Request) bool {return true},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error){
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

// define reader > listen to new msg sent to our WebSocket 
func Reader(conn *websocket.Conn){
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message (for clarity)
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func Writer(conn *websocket.Conn){
	for {
		fmt.Println("sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}