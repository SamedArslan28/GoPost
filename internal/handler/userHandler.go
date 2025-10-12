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

type RegisterForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,min_length=5,email"`
}

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

type EmailRequest struct {
	Email string `json:"email"`
}

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
