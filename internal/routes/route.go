package routes

import (
	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	user := app.Group("/user")
	user.Post("/register", userHandler.RegisterHandler)
	user.Post("/find/email", userHandler.FindByEmailHandler)
}
