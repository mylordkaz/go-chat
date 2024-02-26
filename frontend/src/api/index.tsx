let socket = new WebSocket("ws://localhost:8080/ws")


// connect to the WebSocket endpoint and listens for events
let connect = () =>{
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

// allows to send msg, frontend >>> backend 
let sendMsg = (msg: string) => {
    console.log("sending msg:", msg)
    socket.send(msg)
}

export default {connect, sendMsg}