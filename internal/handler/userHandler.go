package handler

import (
	"database/sql"
	"errors"

	"github.com/SamedArslan28/gopost/internal/models"
	"github.com/SamedArslan28/gopost/internal/service"
	"github.com/SamedArslan28/gopost/internal/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service service.UserService
}
type RegisterForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,min_length=5,email"`
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterHandler(c *fiber.Ctx) error {
	var req RegisterForm
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	newUser, err := h.service.Register(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": newUser,
	})
}

type EmailRequest struct {
	Email string `json:"email"`
}

func (h *UserHandler) FindByEmailHandler(c *fiber.Ctx) error {
	var req EmailRequest
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "failed to parse req",
		})
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	user, err := h.service.GetByEmail(c.Context(), req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get user: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

type IDRequest struct {
	ID int `json:"id"`
}

//func (h *UserHandler) FindByIdHandler(c *fiber.Ctx) error {
//	var req IDRequest
//	err := c.BodyParser(&req)
//	if err != nil {
//		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
//			"error": "failed to parse req",
//		})
//	}
//
//	user, err := h.service.GetById(c.Context(), req.ID)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//				"error": "user not found",
//			})
//		}
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"error": "failed to get user: " + err.Error(),
//		})
//	}
//	return c.Status(fiber.StatusOK).JSON(fiber.Map{
//		"user": user,
//	})
//}
