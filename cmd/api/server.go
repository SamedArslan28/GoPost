package main

import (
	"log"

	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/SamedArslan28/gopost/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Server holds the Fiber application instance and its dependencies.
type Server struct {
	App *fiber.App
}

// NewServer acts as a provider for our Server. It receives the dependencies
// it needs (the userHandler) and sets up the Fiber application.
func NewServer(userHandler handler.UserHandler) *Server {
	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app, &userHandler)

	return &Server{
		App: app,
	}
}

// Start runs the Fiber server and blocks forever.
func (s *Server) Start() {
	log.Println("Starting server on port 3000")
	err := s.App.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
