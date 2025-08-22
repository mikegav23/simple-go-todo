package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-go-todo/internal/db"

	"golang.org/x/crypto/bcrypt"
)

// GET /users -- Get All Users
func (h *Handler) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.Queries.ListUsers(context.Background())
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		fmt.Println("DB error: ", err.Error())
		return
	}
	fmt.Println(users)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

type CreateUserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// POST /users -- Create New User
func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Decode the JSON body
    var req CreateUserRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // 2. Validate input
    if req.Username == "" || req.Password == "" {
        http.Error(w, "Username and password are required", http.StatusBadRequest)
        return
    }

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert user into db and print
	insertedUser, err := h.Queries.CreateUser(context.Background(), db.CreateUserParams{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		fmt.Println("DB error: ", err.Error())
		return
	}
	fmt.Println("Inserted User: ", insertedUser)

	if err := json.NewEncoder(w).Encode(insertedUser); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}