import { useState } from "react";
import "./App.css";
import { Routes, Route } from "react-router-dom";
import LandingPage from "./pages/landingpage";
import VillagePage from "./pages/game/village";
import NotFound from "./pages/notfound";

function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/game/village/:id" element={<VillagePage />} />
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}

export default App;
