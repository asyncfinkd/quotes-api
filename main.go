package main

import (
	"context"
	"log"
	"quotes-api/config"
	"time"

	"quotes-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

// const dbName = "fiber_test"
const mongoURI = "mongodb+srv://giga:vivomini@rest.nl9di.mongodb.net/quotes?retryWrites=true&w=majority"

// Employee struct
type Authors struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Text string `json:"text,omitempty" bson:"text,omitempty"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database("quotes")

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

// @title Quotes App
// @version 1.0
// @description This is an API for Quotes Application

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	// router.SetupUserRoutes(app)
	router.SetupAuthRoutes(app)

	app.Static("/", "./uploads")

	app.Get("/test/db", func(ctx *fiber.Ctx) error {
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("authors").Find(ctx.Context(), query)
		if err != nil {
			return ctx.Status(500).SendString(err.Error())
		}

		var authors []Authors = make([]Authors, 0)

		if err := cursor.All(ctx.Context(), &authors); err != nil {
			return ctx.Status(500).SendString(err.Error())

		}

		return ctx.JSON(authors)
	})

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
