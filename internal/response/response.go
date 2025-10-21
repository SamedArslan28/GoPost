package response

import (
	"errors"
	"log"

	apperrors "github.com/SamedArslan28/gopost/internal/errors"
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents a generic error JSON.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ValidationErrorResponse represents validation errors.
type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// SuccessResponse represents success messages or data.
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// JSONError returns a standard error response with status code.
func JSONError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(ErrorResponse{Error: message})
}

// JSONValidationError returns a standardized validation error response.
func JSONValidationError(c *fiber.Ctx, errs map[string]string) error {
	return c.Status(fiber.StatusBadRequest).JSON(ValidationErrorResponse{Errors: errs})
}

// JSONSuccess returns a standardized success response.
func JSONSuccess(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(SuccessResponse{Data: data})
}

func HandleError(c *fiber.Ctx, err error) error {
	if errors.Is(err, apperrors.ErrNotFound) {
		return JSONError(c, fiber.StatusNotFound, err.Error())
	}
	if errors.Is(err, apperrors.ErrForbidden) {
		return JSONError(c, fiber.StatusForbidden, err.Error())
	}
	if errors.Is(err, apperrors.ErrParseBody) {
		return JSONError(c, fiber.StatusUnprocessableEntity, err.Error())
	}
	if errors.Is(err, apperrors.ErrEmailConflict) {
		return JSONError(c, fiber.StatusConflict, err.Error())
	}
	log.Printf("An unexpected error occurred: %v", err)

	return JSONError(c, fiber.StatusInternalServerError, "Internal Server Error")
}
