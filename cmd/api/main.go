package main

import (
	"log"
	"os"

	"github.com/SamedArslan28/gopost/internal/validator"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load Configuration from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbDsn := os.Getenv("POSTGRES_URL")
	if dbDsn == "" {
		log.Fatal("POSTGRES_URL environment variable not set")
	}

	// 2. Initialize the application using Wire's generated function
	// This single call will create and connect the database, repository,
	// service, handler, and the Fiber app itself.
	server, err := InitializeApp(dbDsn)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	log.Println("Dependencies initialized")

	// 3. Perform any initial setup that isn't part of the dependency graph
	validator.InitValidator()

	// 4. Start the server
	server.Start()
}
