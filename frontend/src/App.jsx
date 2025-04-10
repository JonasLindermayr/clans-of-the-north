import { Routes, Route } from "react-router-dom";
import LandingPage from "./pages/landingpage";
import VillagePage from "./pages/game/village";
import NotFound from "./pages/notfound";
import { Login } from "./pages/login";
import { SignUp } from "./pages/signup";

function App() {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<SignUp />} />
      <Route path="/game/village/:id" element={<VillagePage />} />
      <Route path="*" element={<NotFound />} />
    </Routes>
  );
}

export default App;
