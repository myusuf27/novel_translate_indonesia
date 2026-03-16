package main

import (
	"log"
	"novel_translate_indonesia/internal/database"
	"novel_translate_indonesia/internal/handler"
	"novel_translate_indonesia/internal/repository"
	"novel_translate_indonesia/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize database
	database.InitDB()

	// Initialize Fiber engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Use(logger.New())
	app.Static("/static", "./static")

	// Initialize Repositories and Services
	repo := repository.NewChapterRepository()
	trans := service.NewTranslator()

	// Initialize Handlers
	h := handler.NewChapterHandler(repo, trans)

	// Routes
	app.Get("/", h.Index)
	app.Post("/sync", h.Sync)
	app.Get("/chapter/:id", h.Show)
	app.Post("/chapter/:id/translate", h.Translate)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
