package main

import (
	"fmt"
	"log"
	"net/http"

	
)



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