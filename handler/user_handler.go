package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Quotes struct {
	ID       uint64   `json:"id"`
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
		if _id == v.ID {
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
			if category == t {
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
