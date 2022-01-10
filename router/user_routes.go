package router

import (
	"quotes-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupUserRoutes(app *fiber.App) {
	// GlobalPrefix "/api"
	api := app.Group("/api", logger.New())

	// Quotes
	api.Get("/quotes", handler.GetQuotes)
	api.Get("/quotes/:id", handler.GetOnceQuotes)
	api.Get("/quotes/category/:category", handler.GetOnceQuotesByFilter)

	// Authors
	api.Get("/authors", handler.GetAuthors)
	api.Get("/authors/:id", handler.GetOnceAuthors)
	api.Get("/authors/category/:category", handler.AuthorsFilter)
}
