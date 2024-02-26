import '../chatroom.css';
import { useEffect, useState } from "react";
import { connect, sendMsg } from "../api";


interface Message {
  text: string
  sender: string
}


export default function Chatroom() {
  const [messages, setMessages] = useState<Message[]>([])
  const [newMessage, setNewMessage] = useState<string>("")


  useEffect(() => {
    connect()

    return () => {}
  }, [])

  const send = (e: React.FormEvent) => {
    e.preventDefault()

    if (newMessage.trim() !== ""){
      sendMsg(newMessage)
      setMessages([...messages, {text: newMessage, sender: "me"}])
      setNewMessage("")
    }
    
  }

    return (
      <>
        <div className="chat">
          <div className="chat-container">
            {messages.map((message, index) => (
              <div className='message' key={index}>
                {message.text}
              </div>
            ))}
  
            <form className="form" onSubmit={send} >
              <input 
                placeholder="say something nice"
                value={newMessage}
                onChange={(e) => setNewMessage(e.target.value)} />
              <button type='submit' className="chat-btn">Hit</button>
            </form>
          </div>
        </div>
      </>
    );
  }
  