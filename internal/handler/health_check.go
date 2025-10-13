package handler

import "github.com/gofiber/fiber/v2"

// HealthCheck godoc
// @Summary Show the status of the server
// @Description Returns a simple message to check if the API is alive
// @Tags root
// @Accept json
// @Produce json
// @Success 202 {object} map[string]interface{}
// @Router /healthcheck [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Hello World",
		"status":  fiber.StatusOK,
	})
}
