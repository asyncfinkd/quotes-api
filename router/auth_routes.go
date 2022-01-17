package router

import (
	"quotes-api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupAuthRoutes(app *fiber.App) {
	// GlobalPrefix "/api"
	api := app.Group("/api", logger.New())

	// Auth
	api.Post("/signin", handler.Signin)
	api.Post("/signup", handler.Signup)
}
