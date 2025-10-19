package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	go func() {
		server.Start()
	}()

	gracefulShutdown(server, 10*time.Second)

}

// gracefulShutdown gracefully stops the server with a timeout
func gracefulShutdown(server *Server, timeout time.Duration) {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown signal received. Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- server.Shutdown()
	}()

	select {
	case <-ctx.Done():
		log.Println("⚠️ Timeout reached during shutdown. Forcing exit...")
	case err := <-done:
		if err != nil {
			log.Printf("Error during shutdown: %v", err)
		} else {
			log.Println("✅ Fiber shut down gracefully.")
		}
	}

	fmt.Println("Running cleanup tasks...")

	fmt.Println("Cleanup completed. Exiting.")
}
