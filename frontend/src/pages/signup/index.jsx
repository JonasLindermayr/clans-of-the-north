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
import "./signup.css";
import { useState } from "react";
import axios from "axios";

export function SignUp() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [username, setUsername] = useState("");
  const [error, setError] = useState("");

  const handleSignUp = async (e) => {
    e.preventDefault();
    setError("");

    if (password !== confirmPassword) {
      setError("Passwords do not match");
      return;
    }

    try {
      const response = await axios.post("http://localhost:8080/auth/signup", {
        email,
        password,
        username,
      });
    } catch (err) {
      setError("Signup failed: " + (err.response?.data?.error || err.message));
    }
  };

  return (
    <div className="wrapper">
      <Paper className="form" radius={0} p={30}>
        <Title order={2} className="title" ta="center" mt="md" mb={50}>
          Clans of the North!
        </Title>

        <TextInput
          label="Email address"
          placeholder="hello@gmail.com"
          size="md"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <TextInput
          label="Username"
          placeholder="Your username"
          size="md"
          mt="md"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />

        <PasswordInput
          label="Password"
          placeholder="Your password"
          mt="md"
          size="md"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <PasswordInput
          label="Confirm Password"
          placeholder="Confirm your password"
          mt="md"
          size="md"
          value={confirmPassword}
          onChange={(e) => setConfirmPassword(e.target.value)}
        />

        {error && (
          <Text c="red" mt="md">
            {error}
          </Text>
        )}

        <Checkbox label="Keep me logged in" mt="xl" size="md" />

        <Button fullWidth mt="xl" size="md" onClick={handleSignUp}>
          Sign up!
        </Button>

        <Text ta="center" mt="md">
          Already have an account?{" "}
          <Anchor href="/login" fw={700}>
            Login
          </Anchor>
        </Text>
      </Paper>
    </div>
  );
}
