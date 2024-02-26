package websocket

import "fmt"

//define a pool struct which contain all the channels needed, also a map of clients
type Pool struct{
	Register 	chan *Client
	Unregister 	chan *Client
	Clients  	map[*Client]bool
	Broadcast 	chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan Message),
	 }
}

// constanly listen for anything passed to any pool's chan
func (pool *Pool) Start(){
	for {
		select {
		// send "New User Joined" to all of the clients of the pool
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool:", len(pool.Clients))
			for client := range pool.Clients{
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "new User Joined..."})
			}
			
		// send "user disconnect" to all other clients of the pool
		case client := <-pool.Unregister:
            delete(pool.Clients, client)
            fmt.Println("Size of Connection Pool: ", len(pool.Clients))
            for client := range pool.Clients {
                client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
            }
            
		//when passed a msg loop through all client in the pool and send the msg 
		case message := <-pool.Broadcast:
            fmt.Println("Sending message to all clients in Pool")
            for client := range pool.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    return
                }
            }
		}
	}
}