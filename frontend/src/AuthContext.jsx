import { createContext, useState, useEffect } from "react";
import axios from "axios";

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [token, setToken] = useState(sessionStorage.getItem("token"));

  const login = (newToken, user) => {
    sessionStorage.setItem("token", newToken);
    setUser(user);
    setToken(newToken);
  };

  const logout = () => {
    sessionStorage.removeItem("token");
    setToken(null);
    setUser(null);
  };

  const verifyToken = async () => {
    try {
      const res = await axios.get("/api/auth/verify", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      return true;
    } catch (err) {
      logout();
      return false;
    }
  };

  return (
    <AuthContext.Provider value={{ user, token, login, logout, verifyToken }}>
      {children}
    </AuthContext.Provider>
  );
};
