import '../chatroom.css';
import { useEffect, useState } from "react";
import { connect, sendMsg } from "../api";
import { useLocation } from 'react-router-dom';

 


interface Message {
  type: number
  text: string
  sender: string
  
  }



export default function Chatroom() {
  const [messages, setMessages] = useState<Message[]>([])
  const [newMessage, setNewMessage] = useState<string>("")

  const location = useLocation()
  const user = new URLSearchParams(location.search).get("username") || "guest"

  



  useEffect(() => {
    const handleNewMessage = (event: MessageEvent) => {
       const receivedMessage: Message = JSON.parse(event.data);
      setMessages((prevMessages) => [...prevMessages, receivedMessage]);
      
    };

    connect(handleNewMessage);

    

    return () => {}
  }, [])
 

  const send = (e: React.FormEvent) => {
    e.preventDefault()

    // if (newMessage.trim() !== "" && sendMessage){
    //   sendMessage(JSON.stringify({ text: newMessage, sender: user }))
    //   setMessages([...messages, {data: newMessage}])
    //   setNewMessage("")
    // }
    if (newMessage.trim() !== "") {
      sendMsg(JSON.stringify({ type: 1, text: newMessage, sender: user }));
      setNewMessage("");
    }
  }

    return (
      <>
        <div className="chat">
          <header><h1>Go-chat</h1></header>
          <div className="chat-container">
            {messages.map((message, index) => (
              <div key={index} className={message.sender === user ? 'sent' : 'received'} >
                <strong>{message.sender}</strong>
                <div className='message'>
                {message.text}
                </div>
              </div>
            ))}
  
            <form className="form" onSubmit={send} >
              <input 
                placeholder="say something nice"
                value={newMessage}
                onChange={(e) => setNewMessage(e.target.value)} />
              <button type='submit' className="chat-btn">send</button>
            </form>
          </div>
        </div>
      </>
    );
  }
  