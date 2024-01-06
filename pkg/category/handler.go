package category

import (
	"log"
	"strconv"

	"article-go-api/pkg/util"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Handler struct holds required services for handler to function
type Handler struct {
    service *Service
}

// NewHandler initializes a new category handler
func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

// Validator instance
var validate = validator.New()


// CreateCategory handles POST requests to create a new category
func (h *Handler) CreateCategory(c *fiber.Ctx) error {
    var dto CreateCategoryDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := validate.Struct(dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": util.GetValidationErrors(err)})
    }

    cat := Category{Name: dto.Name}
    // Call service to create the category...
	h.service.CreateCategory(&cat)

    return c.Status(fiber.StatusCreated).JSON(cat)
}

// GetAllCategories handles GET requests to retrieve all categories
func (h *Handler) GetAllCategories(c *fiber.Ctx) error {
    categories, err := h.service.GetAllCategories()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve categories"})
    }

    return c.JSON(categories)
}

// GetCategory handles GET requests to retrieve a single category by ID
func (h *Handler) GetCategory(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    // log id
    log.Println(id)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    category, err := h.service.GetCategory(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
    }

    return c.JSON(category)
}

// UpdateCategory handles PUT requests to update a category
func (h *Handler) UpdateCategory(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var dto UpdateCategoryDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := validate.Struct(dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": util.GetValidationErrors(err)})
    }

    cat := Category{ID: id, Name: dto.Name}
	// Call service to update the category...
	h.service.UpdateCategory(&cat)
    return c.JSON(cat)
}

// DeleteCategory handles DELETE requests to delete a category
func (h *Handler) DeleteCategory(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }

	h.service.DeleteCategory(id)
    return c.SendStatus(fiber.StatusNoContent)
}
