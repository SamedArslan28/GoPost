package handler

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	res := fiber.Map{
		"status":  "ok",
		"message": "Server is up and running!",
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
