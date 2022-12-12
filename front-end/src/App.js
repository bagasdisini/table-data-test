import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import Add from "./pages/Add";
import Detail from "./pages/Detail";
import Edit from "./pages/Edit";
import ShowAll from "./pages/ShowAll";
import {
  BrowserRouter as Router,
  Route,
  Routes,
} from "react-router-dom";
import React from "react";

function App() {
  return (
      <Routes>
        <Route path="/" element={<ShowAll />} />
        <Route path="/add" element={<Add />} />
        <Route path="/detail/:id" element={<Detail />} />
        <Route path="/edit/:id" element={<Edit />} />
      </Routes>
  );
}

export default App;
