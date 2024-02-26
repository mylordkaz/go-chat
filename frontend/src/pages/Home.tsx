export default function Home() {
    return (
      <div className="home">
        <div className="card">
          <h1>Welcome!</h1>
          <div className="start">
            <label>Enter your name:</label>
            <input type="text" />
            <button className="enter-btn">Enter Chat</button>
          </div>
        </div>
      </div>
    );
  }