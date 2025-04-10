import {
  Anchor,
  Button,
  Checkbox,
  Paper,
  PasswordInput,
  Text,
  TextInput,
  Title,
} from "@mantine/core";
import { useState } from "react";
import axios from "axios";
import { jwtDecode } from "jwt-decode";
import "./login.css";

export function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post("http://localhost:8080/auth/login", {
        email,
        password,
      });

      sessionStorage.setItem("token", response.data.token);

      const decoded = jwtDecode(response.data.token);
      console.log(decoded);
    } catch (err) {
      setError("Login failed: " + (err.response?.data?.error || err.message));
    }
  };

  return (
    <div className="wrapper">
      <Paper className="form" radius={0} p={30}>
        <Title order={2} className="title" ta="center" mt="md" mb={50}>
          Clans of the North!
        </Title>

        {error && (
          <Text c="red" ta="center" mb="md">
            {error}
          </Text>
        )}

        <form onSubmit={handleLogin}>
          <TextInput
            label="Email address"
            placeholder="hello@gmail.com"
            size="md"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
          <PasswordInput
            label="Password"
            placeholder="Your password"
            mt="md"
            size="md"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <Checkbox label="Keep me logged in" mt="xl" size="md" />
          <Button fullWidth mt="xl" size="md" type="submit">
            Login
          </Button>
        </form>

        <Text ta="center" mt="md">
          Don&apos;t have an account?{" "}
          <Anchor href="/signup" fw={700}>
            Sign up!
          </Anchor>
        </Text>
      </Paper>
    </div>
  );
}
