package main

import (
	"log"
	"quotes-api/config"
	"quotes-api/database"

	"quotes-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	router.SetupUserRoutes(app)
	router.SetupAuthRoutes(app)

	app.Static("/", "./uploads")

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Route is not found on the server",
		})
	})

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	log.Fatal(app.Listen(config.Config("PORT")))
}
