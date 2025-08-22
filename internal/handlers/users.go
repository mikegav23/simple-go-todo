package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"simple-go-todo/internal/db"
	"strconv"

	"github.com/go-chi/chi/v5"
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

	// 3. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// 4. Insert user into db and print
	insertedUser, err := h.Queries.CreateUser(context.Background(), db.CreateUserParams{
		Username:     string(req.Username),
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

type UpdateUserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// PUT /users/{userID} -- Update User (requires both username and password even if unchanged)
func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Parse userID from URL parameters (which is a string)
	userIDParam := chi.URLParam(r, "userID")

	// 2. Convert userIDParam (string) to int32 using strconv.Atoi
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// 3. Decode the JSON body
    var req UpdateUserRequest
    err = json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // 4. Validate input
    if req.Username == "" || req.Password == "" {
        http.Error(w, "Username and password are required", http.StatusBadRequest)
        return
    }

	// 5. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// 6. Update user credentials
	updatedUser, err := h.Queries.UpdateUser(context.Background(), db.UpdateUserParams{
		ID: 		  int32(userID),
		Username:     string(req.Username),
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		fmt.Println("DB error: ", err.Error())
		return
	}
	fmt.Println("Updated user: ", updatedUser.ID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Credentials updated successfully"})
}

// DELETE /users/{userID} -- Delete User
func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Parse userID from URL parameters (which is a string)
	userIDParam := chi.URLParam(r, "userID")

	// 2. Convert userIDParam (string) to int32 using strconv.Atoi
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// 3. Delete user
	deletedUser, err := h.Queries.DeleteUser(context.Background(), int32(userID))
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		fmt.Println("DB error: ", err.Error())
		return
	}
	fmt.Println("Deleted user: ", deletedUser.ID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}