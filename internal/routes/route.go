package routes

import (
	_ "github.com/SamedArslan28/gopost/docs"
	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {

	app.Get("/swagger/*", swagger.New(swagger.Config{
		Title:       "GoPost API Docs",
		DeepLinking: true,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/", fiber.StatusFound)
	})

	app.Get("healthcheck", handler.HealthCheck)
	user := app.Group("/user")
	user.Post("/register", userHandler.RegisterHandler)
	user.Post("/find/email", userHandler.FindByEmailHandler)
}
