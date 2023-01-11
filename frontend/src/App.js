
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from './pages/Login';
import ChecklistPencarian from "./pages/ChecklistPencarian";
import Home from "./pages/Home";
import './App.css';
import Sidebars from "./components/Sidebar";
import { useEffect, useState } from "react";

function App() {
  const [isLoggedin, setIsLoggedin] = useState();

  useEffect(() => {
    const isLoggedinLS = localStorage.getItem("isLoggedin");
    isLoggedinLS ? setIsLoggedin(true) : setIsLoggedin(false);
  }, []);

  return (
    <BrowserRouter>
      {isLoggedin ? (
        <div className="d-flex">
          <div>
            <Sidebars />
          </div>
          <div className="w-100">
            <Routes>
              <Route exact path="/" element={<Home />} />
              <Route exact path="/checklistpencairan" element={<ChecklistPencarian />} />
              <Route path="*" element={<Home />} />
            </Routes>
          </div>
        </div>) : <div>
        <Routes>
          <Route exact path="/" element={<Login />} />
          <Route path="*" element={<Login />} />
        </Routes>
      </div>}

    </BrowserRouter>
  )
}

export default App;
