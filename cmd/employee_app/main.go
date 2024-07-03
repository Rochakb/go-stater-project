package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rochakb/go-stater-project/internal/api"
	"github.com/Rochakb/go-stater-project/internal/repository"
	"github.com/Rochakb/go-stater-project/internal/service"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Fetch environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct database connection string
	dbURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	//dbURI := fmt.Sprintf("postgresql://root@unbxds-MacBook-Pro.local:26257/emp?sslmode=disable")
	if dbPassword == "" {
		dbURI = fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", dbUser, dbHost, dbPort, dbName)
	}
	// "postgresql://root@unbxds-MacBook-Pro.local:26257/emp?sslmode=disable"
	// Initialize your PostgreSQL database connection
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Initialize repository and service
	repo := repository.NewPostgreSQLRepository(db)
	svc := service.NewEmployeeService(repo)

	// Initialize endpoints
	endpoints := api.NewEndpoints(svc)

	// Start HTTP server
	addr := ":8082"
	fmt.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, api.MakeHTTPHandler(endpoints)); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
