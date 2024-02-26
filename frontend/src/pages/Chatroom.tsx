import '../chatroom.css';
import { useEffect } from "react";
import { connect, sendMsg } from "../api";
export default function Chatroom() {

  useEffect(() => {
    connect()

    return () => {}
  }, [])

  const send = () => {
    console.log("hello")
    sendMsg("hello")
  }

    return (
      <>
        <div className="chat">
          <div className="chat-container">
            <div className="message">bro!!!</div>
            <div className="message">what's up</div>
            <div className="message">call me</div>
            <div className="message">or not... 🤣 </div>
            <div className="message">did you deploy ???</div>
  
            <form className="form">
              <input placeholder="say something nice" />
              <button className="chat-btn" onClick={send} >Hit</button>
            </form>
          </div>
        </div>
      </>
    );
  }
  