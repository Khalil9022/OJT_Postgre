
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from './pages/Login';
import ChecklistPencarian from "./pages/ChecklistPencarian";
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route exact path="/" element={<Login />} />
        <Route exact path="/home" element={<ChecklistPencarian />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App;
