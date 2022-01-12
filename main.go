package main

import (
	"log"
	"quotes-api/config"

	"quotes-api/database"
	"quotes-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// @title Quotes App
// @version 1.0
// @description This is an API for Quotes Application

func main() {
	// Connect with database
	database.ConnectDb()
	// if err := Connect(); err != nil {
	// log.Fatal(err)
	// }

	app := fiber.New()

	app.Use(cors.New())

	// app.Get("/test/db", func(ctx *fiber.Ctx) error {
	// 	rows, err := db.Query("SELECT * FROM quotes")
	// 	if err != nil {
	// 		return ctx.Status(500).SendString(err.Error())
	// 	}
	// 	defer rows.Close()
	// 	result := Todos{}

	// 	for rows.Next() {
	// 		employee := Todo{}
	// 		if err := rows.Scan(&employee.Id, &employee.Text, &employee.Author); err != nil {
	// 			return err // Exit if we get an error
	// 		}

	// 		result.Todos = append(result.Todos, employee)
	// 	}

	// 	return ctx.Status(fiber.StatusOK).JSON(result)
	// })

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
