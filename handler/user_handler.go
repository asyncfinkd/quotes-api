package handler

import (
	"fmt"
	"quotes-api/database"
	models "quotes-api/models/quotes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorGallery struct {
	Url      string   `json:"url"`
	Category []string `json:"category"`
	Name     string   `json:"name"`
}

func GetQuotes(ctx *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := database.Global().Db.Collection("quotes").Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var quotes []models.Quotes = make([]models.Quotes, 0)

	if err := cursor.All(ctx.Context(), &quotes); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    quotes,
	})
}

func GetOnceQuotes(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "not found quote",
		})
	}

	filter := bson.D{{"_id", _id}}

	var result models.Quotes

	if err := database.Global().Db.Collection("quotes").FindOne(ctx.Context(), filter).Decode(&result); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": true,
			"message": "not found quote",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    result,
	})
}

func AddQuotes(ctx *fiber.Ctx) error {
	collection := database.Global().Db.Collection("quotes")
	quote := new(models.Quotes)

	if err := ctx.BodyParser(quote); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	quote.ID = ""

	_, err := collection.InsertOne(ctx.Context(), quote)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "quote successfully added.",
	})
}

func DeleteQuotes(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(
		ctx.Params("id"),
	)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": true,
			"message": "not found id",
		})
	}

	query := bson.D{{Key: "_id", Value: id}}
	result, err := database.Global().Db.Collection("quotes").DeleteOne(ctx.Context(), &query)

	if err != nil {
		return ctx.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return ctx.SendStatus(404)
	}

	return ctx.Status(fiber.StatusOK).JSON("...")
}

func GetAuthors(ctx *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := database.Global().Db.Collection("authors").Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var authors []models.Authors = make([]models.Authors, 0)

	if err := cursor.All(ctx.Context(), &authors); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    authors,
	})
}

func GetOnceAuthors(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "not found author",
		})
	}

	filter := bson.D{{"_id", _id}}

	var result models.Authors

	if err := database.Global().Db.Collection("authors").FindOne(ctx.Context(), filter).Decode(&result); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": true,
			"message": "not found author",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    result,
	})
}

func GetAuthorsByFilter(ctx *fiber.Ctx) error {
	params := ctx.Params("filter")

	query := bson.D{{}}
	cursor, err := database.Global().Db.Collection("authors").Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var authors []models.Authors = make([]models.Authors, 0)

	if err := cursor.All(ctx.Context(), &authors); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var data []models.Authors

	for _, v := range authors {
		for _, t := range v.Category {
			if strings.ToLower(params) == strings.ToLower(t) {
				data = append(data, v)
			}
		}
	}

	fmt.Println(data)

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"item":    data,
	})
}

func GetAuthorsQuotes(ctx *fiber.Ctx) error {
	params := ctx.Params("author")

	query := bson.D{{}}
	cursor, err := database.Global().Db.Collection("quotes").Find(ctx.Context(), query)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	var quotes []models.Quotes = make([]models.Quotes, 0)

	if err := cursor.All(ctx.Context(), &quotes); err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	fmt.Println(quotes)

	var reData []models.Quotes = make([]models.Quotes, 0)

	filterParams := strings.Split(params, "%20")
	joinedParams := strings.Join(filterParams, " ")
	for _, v := range quotes {
		if strings.ToLower(joinedParams) == strings.ToLower(v.Author) {
			reData = append(reData, v)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    reData,
	})
}

func AddAuthor(ctx *fiber.Ctx) error {
	category := ctx.Params("category")
	splitedCategory := strings.Split(category, "&")

	file, err := ctx.FormFile("document")
	if err != nil {
		return err
	}

	nameValue := ctx.FormValue("Name")
	ctx.SaveFile(file, fmt.Sprintf("./uploads/images/%s", file.Filename))

	collection := database.Global().Db.Collection("authors")

	author := &AuthorGallery{
		Name:     nameValue,
		Category: splitedCategory,
		Url:      file.Filename,
	}

	xml, err := collection.InsertOne(ctx.Context(), author)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	fmt.Println(xml)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "author successfuly added.",
	})
}

func DeleteAuthor(ctx *fiber.Ctx) error {
	_id, err := primitive.ObjectIDFromHex(
		ctx.Params("id"),
	)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": true,
			"message": "not found id",
		})
	}

	query := bson.D{{Key: "_id", Value: _id}}
	result, err := database.Global().Db.Collection("authors").DeleteOne(ctx.Context(), &query)

	if err != nil {
		return ctx.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return ctx.SendStatus(404)
	}

	return ctx.Status(fiber.StatusOK).JSON("...")
}
