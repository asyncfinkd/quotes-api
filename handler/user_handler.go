package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Quotes struct {
	ID       uint     `json:"id"`
	Text     string   `json:"text"`
	Author   string   `json:"author"`
	Category []string `json:"category"`
}

type AuthorGallery struct {
	ID       uint     `json:"id"`
	Url      string   `json:"url"`
	Category []string `json:"category"`
	Name     string   `json:"name"`
}

var quotes = []*Quotes{
	{ID: 0, Text: "There is only one thing that makes a dream impossible to achieve: the fear of failure.", Author: "Paulo Coelho", Category: []string{"Motivation"}},
	{ID: 1, Text: "Have no fear of perfection - you'll never reach it.", Author: "Salvador Dalí", Category: []string{"Motivation"}},
}

var authorGallery = []*AuthorGallery{
	{ID: 0, Name: "Salvador Dali", Url: "salvador-dali.jpeg", Category: []string{"Artist"}},
}

// @Summary Get all Quotes
// @Description Get all Quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Success 200 {array} Quotes{}
// @Router /api/quotes [get]
func GetQuotes(ctx *fiber.Ctx) error {
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
	id := ctx.Params("id")

	_id, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for _, v := range quotes {
		if uint(_id) == v.ID {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"item":    *v,
			})
		}
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "something is wrong.",
	})
}

// @Summary Get Quotes by Category
// @Description Get Quotes by Category
// @Tags Quotes
// @Accept json
// @Produce json
// @Success 200 {array} Quotes{}
// @Router /api/quotes/category/{category} [get]
func GetOnceQuotesByFilter(ctx *fiber.Ctx) error {
	category := ctx.Params("category")

	var reQuotes []*Quotes

	for _, v := range quotes {
		for _, t := range v.Category {
			if strings.ToLower(category) == strings.ToLower(t) {
				reQuotes = append(reQuotes, v)
			}
		}
	}

	if len(reQuotes) > 0 {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"item":    reQuotes,
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "category is not defined",
		})
	}
}

func AddQuotes(ctx *fiber.Ctx) error {
	type request struct {
		Text     *string   `json:"text"`
		Author   *string   `json:"author"`
		Category *[]string `json:"category"`
	}

	var body request

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("body parse error")
	}

	newQuote := &Quotes{
		ID:       uint(len(quotes) + 1),
		Text:     *body.Text,
		Author:   *body.Author,
		Category: *body.Category,
	}

	quotes = append(quotes, newQuote)

	return ctx.Status(fiber.StatusOK).JSON(newQuote)
}

// @Summary Delete quote
// @Description Delete quote
// @Tags Authors
// @Accept json
// @Produce json
// @Success 200 {object} Quotes{}
// @Router /api/delete/quote/{id} [delete]
func DeleteQuotes(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	_id, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("params parse error")
	}

	for i, t := range quotes {
		if t.ID == uint(_id) {
			quotes = append(quotes[0:i], quotes[i+1:]...)

			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"message": "Quotes successfuly deleted",
			})
		}
	}

	return ctx.Status(fiber.StatusBadRequest).JSON("soemthing is wrong")
}

// @Summary Get all Authors
// @Description Get all Authors
// @Tags Authors
// @Accept json
// @Produce json
// @Success 200 {array} AuthorGallery{}
// @Router /api/authors [get]
func GetAuthors(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    authorGallery,
	})
}

func GetOnceAuthors(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	_id, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for _, v := range authorGallery {
		if _id == uint64(v.ID) {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"item":    v,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": false,
		"message": "Information not found on server",
	})
}

func AuthorsFilter(ctx *fiber.Ctx) error {
	category := ctx.Params("category")

	var reAuthor []*AuthorGallery

	for _, v := range authorGallery {
		for _, t := range v.Category {
			if strings.ToLower(category) == strings.ToLower(t) {
				reAuthor = append(reAuthor, v)
			}
		}
	}

	if len(reAuthor) > 0 {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"item":    reAuthor,
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": false,
			"message": "category is not defined",
		})
	}
}

func AddAuthor(ctx *fiber.Ctx) error {
	// Get Params "category"
	category := ctx.Params("category")
	splitedCategory := strings.Split(category, "&")

	// File Upload "field -> document"
	file, err := ctx.FormFile("document")

	if err != nil {
		return err
	}

	// Save value to form-data
	nameValue := ctx.FormValue("Name")

	// Save file to root directory:
	ctx.SaveFile(file, fmt.Sprintf("./uploads/images/%s", file.Filename))

	// Fill Array
	author := &AuthorGallery{
		ID:       uint(len(authorGallery) + 1),
		Name:     nameValue,
		Category: splitedCategory,
		Url:      file.Filename,
	}

	authorGallery = append(authorGallery, author)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    author,
	})
}

func DeleteAuthor(ctx *fiber.Ctx) error {
	// Get Params ":id"
	id := ctx.Params("id")
	_id, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("params parse error")
	}

	for i, todo := range authorGallery {
		if todo.ID == uint(_id) {
			authorGallery = append(authorGallery[0:i], authorGallery[i+1:]...)

			// Response looks like {success, message}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"message": "Author successfuly deleted",
			})
		}
	}

	// Server error.
	return ctx.Status(fiber.StatusOK).JSON("something is wrong")
}
