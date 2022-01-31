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

// Create Fake DB for Quotes
// var quotes = []*constant.Quotes{
// 	{ID: 1, Text: "There is only one thing that makes a dream impossible to achieve: the fear of failure.", Author: "Paulo Coelho", Category: []string{"Motivation"}},
// 	{ID: 2, Text: "Have no fear of perfection - you'll never reach it.", Author: "Salvador Dalí", Category: []string{"Motivation"}},
// 	{ID: 3, Text: "Whatever you are, be a good one.", Author: "Abraham Lincoln", Category: []string{"Motivation"}},
// 	{ID: 4, Text: "Make everything as simple as possible, but not simpler.", Author: "Albert Einstein", Category: []string{"Motivation"}},
// 	{ID: 4, Text: "Creativity is intelligence having fun.", Author: "Albert Einstein", Category: []string{"Motivation"}},
// 	{ID: 5, Text: "The best revenge is massive success.", Author: "Frank Sinatra", Category: []string{"Motivation"}},
// 	{ID: 6, Text: "If people never did silly things, nothing intelligent would ever get done.", Author: "Ludwig Wittgenstein", Category: []string{"Motivation"}},
// 	{ID: 7, Text: "The two most important days in your life are the day you are born and the day you find out why.", Author: "Mark Twain", Category: []string{"Motivation"}},
// 	{ID: 8, Text: "Pain is inevitable. Suffering is optional.", Author: "Haruki Murakami", Category: []string{"Motivation"}},
// 	{ID: 9, Text: "The only truly secure system is one that is powered off, cast in a block of concrete and sealed in a lead-lined room with armed guards.", Author: "Gene Spafford", Category: []string{"Cyber Security"}},
// 	{ID: 10, Text: "The more you know, the more you realize you know nothing.", Author: "Socrates", Category: []string{"Motivation"}},
// 	{ID: 11, Text: "If you reveal your secrets to the wind, you should not blame the wind for revealing them to the trees.", Author: "Kahlil Gibran", Category: []string{"Motivation"}},
// 	{ID: 12, Text: "The best way out is always through.", Author: "Robert Frost", Category: []string{"Motivation"}},
// 	{ID: 13, Text: "If you don't risk anything, you risk even more.", Author: "Erica Jong", Category: []string{"Motivation"}},
// 	{ID: 14, Text: "There are far, far better things ahead than any we leave behind.", Author: "C.S. Lewis", Category: []string{"Motivation"}},
// 	{ID: 15, Text: "Follow your inner moonlight. Don't hide the madness.", Author: "Allen Ginsberg", Category: []string{"Motivation"}},
// 	{ID: 16, Text: "The harder the conflict, the more glorious the triumph.", Author: "Thomas Paine", Category: []string{"Motivation"}},
// 	{ID: 17, Text: "Every strike brings me closer to the next home run.", Author: "Babe Ruth", Category: []string{"Motivation"}},
// 	{ID: 18, Text: "Everybody has talent, but ability takes hard work.", Author: "Michael Jordan", Category: []string{"Motivation"}},
// }

type AuthorGallery struct {
	Url      string   `json:"url"`
	Category []string `json:"category"`
	Name     string   `json:"name"`
}

// @Summary Get all Quotes
// @Description Get all Quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Success 200 {array} Quotes{}
// @Router /api/quotes [get]
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

// @Summary Get once Quotes
// @Description Get once Quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Success 200 {object} Quotes{}
// @Router /api/quotes/{id} [get]
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

// @Summary Delete quote
// @Description Delete quote
// @Tags Authors
// @Accept json
// @Produce json
// @Success 200 {object} Quotes{}
// @Router /api/delete/quote/{id} [delete]
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

// @Summary Get all Authors
// @Description Get all Authors
// @Tags Authors
// @Accept json
// @Produce json
// @Success 200 {array} AuthorGallery{}
// @Router /api/authors [get]
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
