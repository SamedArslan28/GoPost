package main

import (
	"log"
	"os"

	"github.com/SamedArslan28/gopost/internal/database"
	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/SamedArslan28/gopost/internal/repository"
	"github.com/SamedArslan28/gopost/internal/routes"
	"github.com/SamedArslan28/gopost/internal/service"
	"github.com/SamedArslan28/gopost/internal/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDsn := os.Getenv("POSTGRES_URL")

	db, err := database.ConnectDB(dbDsn)
	if err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
	}
	log.Println("Connected to database")
	validator.InitValidator()

	app := fiber.New()
	app.Use(logger.New())

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(*userService)

	routes.SetupRoutes(app, userHandler)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
