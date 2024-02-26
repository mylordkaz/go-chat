package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mylordkaz/go-chat/pkg/websocket"
)

// define WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	
	//upgrade http conn to Ws conn
	ws, err := websocket.Upgrade(w, r)
	if err != nil{
		fmt.Fprintf(w, "%+V\n", err)
	}

	// listen indefinitely for new msg coming
	go websocket.Writer(ws)
	websocket.Reader(ws)
}


// server setup
func setupRoutes() {
	
	// mape `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main(){
	setupRoutes()

	err:= http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("listenAndServe:", err)
	}
	fmt.Println("Distributed Go Chat v0.01")
}