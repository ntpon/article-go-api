package article

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler *Handler) {
    // Prefix for article routes
    router := app.Group("/articles")

    // Define routes
    router.Get("/", handler.GetAllArticles)
    router.Get("/:id", handler.GetArticle)
    router.Post("/", handler.CreateArticle)
    router.Put("/:id", handler.UpdateArticle)
    router.Delete("/:id", handler.DeleteArticle)
}