package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/mylordkaz/go-chat/pkg/websocket"
)

// define WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}


// server setup
func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main(){
	setupRoutes()

	err:= http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("listenAndServe:", err)
	}
	fmt.Println("Distributed Go Chat v0.01")
}