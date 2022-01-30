package router

import (
	"fmt"
	"quotes-api/handler"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupUserRoutes(app *fiber.App) {
	// GlobalPrefix "/api"
	api := app.Group("/api", logger.New())

	// Quotes
	api.Get("/quotes", handler.GetQuotes)
	api.Get("/quotes/:id", handler.GetOnceQuotes)
	// api.Get("/quotes/category/:category", handler.GetOnceQuotesByFilter)

	// Quotes TODO
	api.Post("/add/quote", handler.AddQuotes)
	api.Delete("/delete/quote/:id", handler.DeleteQuotes)
	// api.Patch("/edit/quote/:id", handler.UpdateQuotes)

	// Authors
	api.Get("/authors", handler.GetAuthors)
	api.Get("/authors/:filter", handler.GetAuthorsByFilter)
	api.Get("/authors/:id", handler.GetOnceAuthors)
	// api.Get("/authors/category/:category", handler.AuthorsFilter)

	// Authors TODO
	api.Post("/add/author/:category", handler.AddAuthor)
	api.Delete("/delete/author/:id", handler.DeleteAuthor)

	api.Get("/test", func(ctx *fiber.Ctx) error {
		type Test struct {
			ID        uint   `json:"id" validate:"required,omitempty"`
			Firstname string `json:"firstname" validate:"required"`
			Password  string `json:"password" validate:"gte=10"`
		}

		user := Test{
			ID:        1,
			Firstname: "Mark",
			Password:  "123123k21o3k12o3",
		}

		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			fmt.Println(err.Error())
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON("...")
	})
}
