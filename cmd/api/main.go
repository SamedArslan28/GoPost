package main

import (
	"log"

	"github.com/SamedArslan28/gopost/internal/config"
	"github.com/SamedArslan28/gopost/internal/validator"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error: Failed to load configuration: %v", err)
	}

	server, err := InitializeApp(cfg)
	if err != nil {
		log.Fatalf("Error: Failed to initialize application: %v", err)
	}
	log.Println("Dependencies initialized successfully")

	validator.InitValidator()

	server.Start()
}
