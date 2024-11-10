package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shuxinqiao/event-attendance-tracker/backend/utils"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Handler struct {
	DB *sql.DB
}

// RegisterHandler registers a new user
func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve user claims from the context
	// claims, ok := r.Context().Value("claims").(*utils.Claims)
	// if !ok {
	// 	http.Error(w, "User not authenticated", http.StatusUnauthorized)
	// 	return
	// }

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Insert user into database
	_, err = h.DB.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", user.Username, hashedPassword, user.Role)
	if err != nil {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User registered successfully")
}

// LoginHandler authentication a user and issue a JWT
func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var hashedPassword string
	var role string
	err := h.DB.QueryRow("SELECT password, role FROM users WHERE username=$1", user.Username).Scan(&hashedPassword, &role)
	if err != nil || !utils.CheckPasswordHash(user.Password, hashedPassword) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username, role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	fmt.Fprintf(w, "Login successful")
}
