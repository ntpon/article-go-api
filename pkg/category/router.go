package category

import (
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App, handler *Handler) {
    // Prefix for category routes
    router := app.Group("/categories")

    // Define routes
    router.Get("/", handler.GetAllCategories)
    router.Get("/:id", handler.GetCategory)
    router.Post("/", handler.CreateCategory)
    router.Put("/:id", handler.UpdateCategory)
    router.Delete("/:id", handler.DeleteCategory ) 
}  