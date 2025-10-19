package handler

import (
	"strconv"

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

func (handler *PostHandler) CreatePost(c *fiber.Ctx) error {
	var req NewPostRequest
	userID := c.Locals("user_id").(int32)

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	newPost := models.Post{
		Title: req.Title,
		Body:  req.Body,
	}

	post, err := handler.service.CreatePost(c.Context(), newPost, userID)
	if err != nil {
		return err
	}
	return response.JSONSuccess(c, 200, post)
}

func (handler *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := handler.service.GetAllPostForUser(c.Context(), c.Locals("user_id").(int32))
	if err != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.JSONSuccess(c, 200, posts)
}

func (handler *PostHandler) GetPostById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postID64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid post ID",
		})
	}

	postID := int32(postID64)
	post, err := handler.service.GetPostById(c.Context(), postID)
	if err != nil {
		return response.JSONError(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.JSONSuccess(c, 200, post)
}
