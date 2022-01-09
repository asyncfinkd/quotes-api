package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Quotes struct {
	ID       int      `json:"id"`
	Text     string   `json:"text"`
	Author   string   `json:"author"`
	Category []string `json:"category"`
}

var quotes = []*Quotes{
	{ID: 0, Text: "There is only one thing that makes a dream impossible to achieve: the fear of failure.", Author: "Paulo Coelho", Category: []string{"Motivation"}},
	{ID: 1, Text: "Have no fear of perfection - you'll never reach it.", Author: "Salvador Dalí", Category: []string{"Motivation"}},
}

// @Summary Get once Quotes
// @Description Get once Quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Success 200 {object} Quotes{}
// @Router /api/quotes/{id} [get]
func GetQuotes(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(quotes)
}

func GetOnceQuotes(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	_id, err := strconv.Atoi(id)

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
