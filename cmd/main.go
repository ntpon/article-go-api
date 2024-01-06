package main

import (
	"article-go-api/internal/database"
	"article-go-api/pkg/article"
	"article-go-api/pkg/category"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
    // Initialize Fiber app
    app := fiber.New()
  
    // Initialize database connection
    db := database.InitDB()
    // Print 
    log.Print("Database connection initialized...")
    // Set up services and handlers for categories
    categoryService := category.NewService(db)
    categoryHandler := category.NewHandler(categoryService)
    category.SetupRoutes(app, categoryHandler)

    // Set up services and handlers for articles
    articleService := article.NewService(db)
    articleHandler := article.NewHandler(articleService)
    article.SetupRoutes(app, articleHandler)

    // Start the server
    err := app.Listen(":3003")
    if err != nil {
        panic(err)
    }
}