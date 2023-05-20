package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var postgres *sql.DB

func setupDatabase() {
	connectionString := fmt.Sprintf(DatabaseConnectionString, getPostgresPass(), getPostgresHost())

	_, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to the DB with error: %v", err)
	}
}

func getPostgresUser() string {
	postgresUser := os.Getenv("POSTGRE_USER")
	if postgresUser == "" {
		return "postgres"
	}
	return postgresUser
}

func getPostgresPass() string {
	postgresPass := os.Getenv("POSTGRES_PASSWORD")
	if postgresPass == "" {
		return "password"
	}
	return postgresPass
}

func getPostgresHost() string {
	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		return "localhost"
	}
	return postgresHost
}

func GetDatabase() *sql.DB {
	return postgres
}
