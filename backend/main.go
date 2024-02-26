package main

import (
	"fmt"
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

// define reader > listen to new msg sent to our WebSocket 
func reader(conn *websocket.Conn){
	for {
	// read in a message	
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

// define WebSocket endpoint 
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	//upgrade http conn to Ws conn
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}

	// listen indefinitely for new msg coming
	reader(ws)
}


// server setup
func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	// mape `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main(){
	setupRoutes()

	err:= http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("listenAndServe:", err)
	}
}