package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Database connection details
const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "ttofis123"
	dbName     = "postgres"
)

// User represents a user entity in the database
type User struct {
	ID       int
	Username string
	Password string
}

var db *sql.DB

func main() {
	// Establish the database connection
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Define the HTTP endpoints
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		// Query the database for the user
		user, err := getUser(username)
		if err != nil {
			http.Error(w, "Login failed", http.StatusInternalServerError)
			return
		}

		// Check if the password is correct
		if user != nil && user.Password == password {
			fmt.Fprint(w, "Login successful")
		} else {
			fmt.Fprint(w, "Invalid username or password")
		}
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		// Create a new user in the database
		err := createUser(username, password)
		if err != nil {
			http.Error(w, "Signup failed", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Signup successful")
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")

		// Delete the user from the database
		err := deleteUser(username)
		if err != nil {
			http.Error(w, "Delete failed", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Delete successful")
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.FormValue("username")
		newPassword := r.FormValue("new_password")

		// Update the user's password in the database
		err := updateUser(username, newPassword)
		if err != nil {
			http.Error(w, "Update failed", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "Update successful")
	})

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUser(username string) (*User, error) {
	query := "SELECT id, username, password FROM users WHERE username=$1"
	row := db.QueryRow(query, username)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func createUser(username, password string) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := db.Exec(query, username, password)
	if err != nil {
		return err
	}

	return nil
}

func deleteUser(username string) error {
	query := "DELETE FROM users WHERE username=$1"
	_, err := db.Exec(query, username)
	if err != nil {
		return err
	}

	return nil
}

func updateUser(username, newPassword string) error {
	query := "UPDATE users SET password=$1 WHERE username=$2"
	_, err := db.Exec(query, newPassword, username)
	if err != nil {
		return err
	}

	return nil
}
