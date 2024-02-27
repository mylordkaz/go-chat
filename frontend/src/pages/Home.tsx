import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Home() {
  const navigate = useNavigate()
  const [username, setUsername] = useState<string>("")

  const handleEnterChat = () => {
    navigate(`/chat?username=${encodeURIComponent(username)}`)
  }

    return (
      <div className="home">
        <header><h1>Go-chat</h1></header>
        <div className="card">
          <h1>Welcome!</h1>
          <div className="start">
            <label>Enter your name:</label>
            <input 
              type="text"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
             />
            <button className="enter-btn" onClick={handleEnterChat}>
              Enter Chat
            </button>
          </div>
        </div>
      </div>
    );
  }