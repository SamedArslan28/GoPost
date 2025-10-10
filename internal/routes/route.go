package routes

import (
	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {

	api := app.Group("/")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message": "Hello World",
			"status":  fiber.StatusOK,
		})
	})

	user := api.Group("/user")
	user.Post("/register", userHandler.RegisterHandler)
	user.Post("/find/email", userHandler.FindByEmailHandler)
	user.Post("/find/id", userHandler.FindByIdHandler)
}
