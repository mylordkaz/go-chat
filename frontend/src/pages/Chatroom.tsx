import '../chatroom.css';
import { useEffect, useState } from "react";
import { connect, sendMsg } from "../api";
import { useLocation } from 'react-router-dom';


interface Message {
  text: string
  sender: string
}


export default function Chatroom() {
  const [messages, setMessages] = useState<Message[]>([])
  const [newMessage, setNewMessage] = useState<string>("")

  const location = useLocation()
  const user = new URLSearchParams(location.search).get("username") || "guest"



  useEffect(() => {
    connect()
    

    return () => {}
  }, [])
  

  const send = (e: React.FormEvent) => {
    e.preventDefault()

    if (newMessage.trim() !== ""){
      sendMsg(newMessage)
      setMessages([...messages, {text: newMessage, sender: user}])
      setNewMessage("")
    }
    
  }

    return (
      <>
        <div className="chat">
          <div className="chat-container">
            {messages.map((message, index) => (
              <div>
                <strong>{message.sender}</strong>
                <div className='message' key={index}>
                {message.text}
                </div>
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
  