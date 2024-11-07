package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/shuxinqiao/event-attendance-tracker/backend/utils"
)

var db *sql.DB

func initDB() {
	var err error

	// Use environment variables for database connection details
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbName)

	for i := 0; i < 5; i++ { // Retry 5 times with a delay
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			if err := db.Ping(); err == nil {
				fmt.Println("Database connected!")
				return
			}
		}
		log.Printf("Database connection attempt %d failed; retrying...\n", i+1)
		time.Sleep(5 * time.Second)
	}

	log.Fatalf("Failed to connect to the database after multiple attempts: %v", err)
	return
}

// createSuperAdmin checks if a super admin exists and creates one if necessary
func createSuperAdmin() {
	username := os.Getenv("SUPERADMIN_USERNAME")
	password := os.Getenv("SUPERADMIN_PASSWORD")
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	// Check if super admin already exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking for super admin: %v", err)
	}

	if !exists {
		// Insert super admin account with hashed password
		_, err = db.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", username, hashedPassword, "super_admin")
		if err != nil {
			log.Fatalf("Error creating super admin: %v", err)
		}
		fmt.Println("Super admin account created successfully.")
	} else {
		fmt.Println("Super admin account already exists.")
	}
}
