
import { useEffect } from 'react'
import './App.css'
import { connect, sendMsg } from "./api";


function App() {

  useEffect(() =>{
    connect()

    return () => {

    }
  }, [])
 

  const send = () => {
    console.log("hello")
    sendMsg("hello")
    
  }

  return (
    <>
    <div className='App'>
      <button onClick={send}>Hit</button>
    </div>
      
    </>
  )
}

export default App
