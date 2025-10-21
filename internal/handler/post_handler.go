package handler

import (
	"errors"
	"strconv"

	apperrors "github.com/SamedArslan28/gopost/internal/errors"
	"github.com/SamedArslan28/gopost/internal/models"
	"github.com/SamedArslan28/gopost/internal/response"
	"github.com/SamedArslan28/gopost/internal/service"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

type NewPostRequest struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var req NewPostRequest
	userID := c.Locals("user_id").(int32)

	if err := c.BodyParser(&req); err != nil {
		return response.JSONError(c, fiber.StatusBadRequest, apperrors.ErrParseBody.Error())
	}

	newPost := models.Post{
		Title: req.Title,
		Body:  req.Body,
	}

	post, err := h.service.CreatePost(c.Context(), newPost, userID)
	if err != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.JSONSuccess(c, 200, post)
}

func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.service.GetAllPostForUser(c.Context(), c.Locals("user_id").(int32))
	if err != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.JSONSuccess(c, 200, posts)
}

func (h *PostHandler) GetPostById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postID64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid post ID",
		})
	}

	postID := int32(postID64)
	post, err := h.service.GetPostById(c.Context(), postID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return response.JSONError(c, fiber.StatusNotFound, err.Error())
		}
		return response.JSONError(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.JSONSuccess(c, 200, post)
}

func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postID64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return response.JSONError(c, fiber.StatusBadRequest, "Invalid post ID format")
	}
	postID := int32(postID64)

	userLocal := c.Locals("user_id")
	if userLocal == nil {
		return response.JSONError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, ok := userLocal.(int32)
	if !ok {
		return response.JSONError(c, fiber.StatusInternalServerError, "Invalid user ID in context")
	}

	err = h.service.DeletePost(c.Context(), postID, userID)
	if err != nil {
		return response.HandleError(c, err)
	}

	return response.JSONSuccess(c, 200, "Delete post successfully")
}

type UpdatePostRequest struct {
	Id    int32  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	var req UpdatePostRequest

	err := c.BodyParser(&req)
	if err != nil {
		return response.JSONError(c, fiber.StatusUnprocessableEntity, apperrors.ErrParseBody.Error())
	}

	userLocal := c.Locals("user_id")
	if userLocal == nil {
		return response.JSONError(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	userID, ok := userLocal.(int32)
	if !ok {
		return response.JSONError(c, fiber.StatusInternalServerError, "Failed to parse user ID from context")
	}

	updatedPost, err := h.service.UpdatePost(c.Context(), req.Id, req.Title, req.Body, userID)
	if err != nil {
		return response.HandleError(c, err)
	}

	return response.JSONSuccess(c, 200, updatedPost)
}
