package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user User
	var dbPassword string

	// Query user and password from the database
	err = db.QueryRow("SELECT id, username, password FROM users WHERE username=$1", creds.Username).
		Scan(&user.ID, &user.Username, &dbPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Check if password matches
	if creds.Password != dbPassword {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Send user information as response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Event struct represents a single event
type Event struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Date     string `json:"date"`
	Location string `json:"location"`
}

// eventsHandler fetches and returns a list of events
func eventsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, date, location FROM events")
	if err != nil {
		http.Error(w, "Unable to retrieve events", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Date, &event.Location)
		if err != nil {
			http.Error(w, "Error scanning events", http.StatusInternalServerError)
			return
		}
		events = append(events, event)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// checkInHandler processes check-in requests
func checkInHandler(w http.ResponseWriter, r *http.Request) {
	var checkInData struct {
		UserID  int `json:"userId"`
		EventID int `json:"eventId"`
	}

	// Decode the JSON body into checkInData
	err := json.NewDecoder(r.Body).Decode(&checkInData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert the check-in record into the checkins table
	_, err = db.Exec("INSERT INTO checkins (user_id, event_id) VALUES ($1, $2)", checkInData.UserID, checkInData.EventID)
	if err != nil {
		http.Error(w, "Failed to check in", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Check-in successful"))
}

func main() {
	initDB()

	http.Handle("/login", corsMiddleware(http.HandlerFunc(loginHandler)))
	http.Handle("/events", corsMiddleware(http.HandlerFunc(eventsHandler)))
	http.Handle("/checkin", corsMiddleware(http.HandlerFunc(checkInHandler)))

	http.Handle("/", corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Event Attendance System!")
	})))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
