package websocket

import (
	
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
