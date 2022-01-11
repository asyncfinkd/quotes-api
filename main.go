package main

import (
	"database/sql"
	"fmt"
	"log"
	"quotes-api/config"
	"quotes-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database instance
var db *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "nikashamiladze"
	password = "none"
	dbname   = "quotes"
)

// Connect function
func Connect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

type Todo struct {
	Todo_id     string `json:"todo_id"`
	Description string `json:"description"`
}

// Employees struct
type Todos struct {
	Todos []Todo `json:"employees"`
}

// @title Quotes App
// @version 1.0
// @description This is an API for Quotes Application

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/test/db", func(ctx *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM todo")
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}
		defer rows.Close()
		result := Todos{}

		for rows.Next() {
			employee := Todo{}
			if err := rows.Scan(&employee.Todo_id, &employee.Description); err != nil {
				return err // Exit if we get an error
			}

			result.Todos = append(result.Todos, employee)
		}

		fmt.Println(rows)
		return ctx.Status(fiber.StatusOK).JSON(&rows)
	})

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
