let socket = new WebSocket("ws://localhost:8080/ws")


// connect to the WebSocket endpoint and listens for events
export const connect = () =>{
    console.log("Attempting Connection...")
    
    socket.onopen = () => {
        console.log("Successfully Connected")
    }

    socket.onmessage = (msg) => {
        
        console.log(msg)
    }

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event)
    }

    socket.onerror = error => {
        console.log("Socket Error: ", error)
    }
}

//allows to send msg, frontend >>> backend 
export const sendMsg = (msg: string) => {
    console.log("sending msg:", msg)
    socket.send(msg)
}
// export const sendMsg = (username: string, message: string) => {
//     const payload = JSON.stringify({ username, message });
//     console.log("sending msg:", payload);
//     socket.send(payload);
//   };
