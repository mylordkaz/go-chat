import './App.css'
import { RouterProvider, createBrowserRouter } from 'react-router-dom';
import Home from './pages/Home';
import Chatroom from './pages/Chatroom';


let router = createBrowserRouter([
  {
    path: "/",
    element: <Home />
  },
  {
    path: "/chat",
    element: <Chatroom />
  }
])

function App() {


  return (
    <div>
      <RouterProvider router={router} />
    </div>
  )
}

export default App
