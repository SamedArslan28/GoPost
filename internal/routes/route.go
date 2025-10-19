package routes

import (
	"github.com/SamedArslan28/gopost/internal/handler"
	"github.com/SamedArslan28/gopost/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler, postHandler *handler.PostHandler) {

	app.Use(middleware.CorsConfig())
	app.Use(middleware.SecurityHeaders())
	//app.Use(middleware.XSSEscapeMiddleware())

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	user := app.Group("/user")
	user.Post("/register", userHandler.RegisterHandler)
	user.Post("/find/email", userHandler.FindByEmailHandler)
	user.Post("/login", userHandler.LoginHandler)

	authenticated := app.Group("/posts", middleware.JWTMiddleware())
	authenticated.Post("/create", postHandler.CreatePost)
	authenticated.Get("/", postHandler.GetAllPosts)
	authenticated.Get("/:id", postHandler.GetPostById)
}
