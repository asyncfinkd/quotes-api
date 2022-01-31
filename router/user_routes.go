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

	// Quotes TODO
	api.Post("/add/quote", handler.AddQuotes)
	api.Delete("/delete/quote/:id", handler.DeleteQuotes)

	// Authors
	api.Get("/authors", handler.GetAuthors)
	api.Get("/authors/:filter", handler.GetAuthorsByFilter)
	api.Get("/authors/:author/quotes", handler.GetAuthorsQuotes)
	api.Get("/authors/:id", handler.GetOnceAuthors)

	// Authors TODO
	api.Post("/add/author/:category", handler.AddAuthor)
	api.Delete("/delete/author/:id", handler.DeleteAuthor)
}
