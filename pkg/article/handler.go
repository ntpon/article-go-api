package article

import (
	"article-go-api/pkg/util"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
    service *Service
}

var validate = validator.New()

func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}


func (h *Handler) CreateArticle(c *fiber.Ctx) error {
    var dto CreateArticleDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := validate.Struct(dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": util.GetValidationErrors(err)})
    }

    article := Article{
        CategoryID: dto.CategoryID,
        Title:      dto.Title,
        Content:    dto.Content,
    }
    
    return c.Status(fiber.StatusCreated).JSON(article)
}

func (h *Handler) GetAllArticles(c *fiber.Ctx) error {
    articles, err := h.service.GetAllArticles()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve articles"})
    }

    return c.JSON(articles)
}

func (h *Handler) GetArticle(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    article, err := h.service.GetArticle(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
    }

    return c.JSON(article)
}

func (h *Handler) UpdateArticle(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var dto UpdateArticleDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := validate.Struct(dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": util.GetValidationErrors(err)})
    }

    article := Article{
        ID:         id,
        CategoryID: dto.CategoryID,
        Title:      dto.Title,
        Content:    dto.Content,
    }

    h.service.UpdateArticle(&article)
    return c.JSON(article)
}

func (h *Handler) DeleteArticle(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    h.service.DeleteArticle(id)
    return c.SendStatus(fiber.StatusNoContent)
}
