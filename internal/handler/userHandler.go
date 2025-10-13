package handler

import (
	"database/sql"
	"errors"

	apperrors "github.com/SamedArslan28/gopost/internal/errors"
	"github.com/SamedArslan28/gopost/internal/models"
	"github.com/SamedArslan28/gopost/internal/response"
	"github.com/SamedArslan28/gopost/internal/service"
	"github.com/SamedArslan28/gopost/internal/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// RegisterForm represents the user registration request body.
// swagger:model
type RegisterForm struct {
	Username string `json:"username" validate:"required" example:"johndoe"`
	Password string `json:"password" validate:"required" example:"secret123"`
	Email    string `json:"email" validate:"required,min_length=5,email" example:"johndoe@example.com"`
}

// EmailRequest represents the request body for searching by email.
// swagger:model
type EmailRequest struct {
	Email string `json:"email" validate:"required,email" example:"johndoe@example.com"`
}

// RegisterHandler godoc
// @Summary Register a new user
// @Description Creates a new user account by providing username, password, and email.
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterForm true "User registration form"
// @Success 201 {object} models.User
// @Failure 400 {object} response.ErrorResponse
// @Failure 422 {object} response.ValidationErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/register [post]
func (h *UserHandler) RegisterHandler(c *fiber.Ctx) error {
	var req RegisterForm
	if err := c.BodyParser(&req); err != nil {
		return response.JSONError(c, fiber.StatusUnprocessableEntity, apperrors.ErrParseBody.Error())
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		return response.JSONValidationError(c, errs)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, "failed to hash password")
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	newUser, err := h.service.Register(c.Context(), user)
	if err != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, "failed to create user: "+err.Error())
	}

	return response.JSONSuccess(c, fiber.StatusCreated, newUser)
}

// FindByEmailHandler godoc
// @Summary Find user by email
// @Description Retrieves a user record using their email address.
// @Tags users
// @Accept json
// @Produce json
// @Param email body EmailRequest true "Email search request"
// @Success 200 {object} models.User
// @Failure 404 {object} response.ErrorResponse
// @Failure 422 {object} response.ValidationErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/find [post]
func (h *UserHandler) FindByEmailHandler(c *fiber.Ctx) error {
	var req EmailRequest
	err := c.BodyParser(&req)
	if err != nil {
		return response.JSONError(c, fiber.StatusUnprocessableEntity, apperrors.ErrParseBody.Error())
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		return response.JSONValidationError(c, errs)
	}

	user, err := h.service.GetByEmail(c.Context(), req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.JSONError(c, fiber.StatusNotFound, "email not found")
		}
		return response.JSONError(c, fiber.StatusInternalServerError, "failed to get user")
	}
	return response.JSONSuccess(c, fiber.StatusOK, user)
}
