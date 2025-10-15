package middleware

import (
	"strings"

	"github.com/SamedArslan28/gopost/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/microcosm-cc/bluemonday"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing Authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format",
			})
		}

		tokenString := parts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}
		if userID, ok := claims["user_id"]; ok {
			c.Locals("user_id", userID)
		}
		return c.Next()
	}
}

func SecurityHeaders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "DENY")
		c.Set("Referrer-Policy", "no-referrer")
		c.Set("Permissions-Policy", "geolocation=(), microphone=()")
		c.Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; sandbox;")
		return c.Next()
	}
}

func CorsConfig() fiber.Handler {
	return cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	})
}

func XSSEscapeMiddleware() fiber.Handler {
	p := bluemonday.UGCPolicy()

	return func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPut || c.Method() == fiber.MethodPatch {
			body := c.Body()
			if len(body) > 0 {
				safeBody := p.SanitizeBytes(body)
				c.Request().SetBody(safeBody)
			}
		}
		return c.Next()
	}
}
