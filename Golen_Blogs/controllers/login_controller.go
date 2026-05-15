package controllers

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	// ✅ Handle preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only POST allowed
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(LoginResponse{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	// 🔐 Hardcoded credentials
	if req.Username == "admin" && req.Password == "1234" {
		json.NewEncoder(w).Encode(LoginResponse{
			Success: true,
			Message: "Login successful",
		})
		return
	}

	// ❌ Invalid login
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(LoginResponse{
		Success: false,
		Message: "Invalid username or password",
	})
}