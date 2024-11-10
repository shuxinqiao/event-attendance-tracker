package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shuxinqiao/event-attendance-tracker/backend/handlers"
	"github.com/shuxinqiao/event-attendance-tracker/backend/middleware"
	"github.com/shuxinqiao/event-attendance-tracker/backend/utils"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
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
	// Database Setup
	initDB()
	createSuperAdmin()

	// Generate a 32-byte random key for JWT signing
	jwtKey, err := utils.GenerateRandomKey(32)
	if err != nil {
		log.Fatal("Failed to generate JWT key:", err)
	}

	// Initialize the utils package with this key
	utils.InitializeJWTKey(jwtKey)

	h := &handlers.Handler{DB: db}

	http.Handle("/register", middleware.CORS(middleware.AuthMiddleware(http.HandlerFunc(h.RegisterHandler))))
	http.Handle("/login", middleware.CORS(http.HandlerFunc(h.LoginHandler)))

	// http.Handle("/checkin", corsMiddleware(http.HandlerFunc(checkInHandler)))

	// Protected route example
	// http.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(ProtectedHandler)))

	http.Handle("/", middleware.CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Event Attendance System!")
	})))

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
