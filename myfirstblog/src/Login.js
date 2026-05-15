import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

function Login() {
  const [user, setUser] = useState({
    username: "",
    password: ""
  });

  const navigate = useNavigate();

  // handle input change
  const handleChange = (e) => {
    setUser({
      ...user,
      [e.target.name]: e.target.value
    });
  };

  // login api
  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:30002/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(user)
      });

      const data = await res.json();

      if (data.success) {
        localStorage.setItem("isLoggedIn", "true");
        localStorage.setItem("user", JSON.stringify(data.user));

        navigate("/blogs");
      } else {
        alert(data.message || "Invalid credentials");
      }

    } catch (err) {
      console.error(err);
      alert("Server error");
    }
  };

  return (
    <div className="container mt-5">
      <div className="card p-4 shadow">
        <h3>Login</h3>

        <form onSubmit={handleLogin}>
          <input
            className="form-control mb-3"
            name="username"
            placeholder="Username"
            value={user.username}
            onChange={handleChange}
          />

          <input
            type="password"
            className="form-control mb-3"
            name="password"
            placeholder="Password"
            value={user.password}
            onChange={handleChange}
          />

          <button className="btn btn-primary w-100">
            Login
          </button>
        </form>
      </div>
    </div>
  );
}

export default Login;