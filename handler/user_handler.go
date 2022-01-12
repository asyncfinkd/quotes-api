package handler

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"quotes-api/constant"
	"quotes-api/database"
	"quotes-api/helper"
	"quotes-api/models"

	"github.com/gofiber/fiber/v2"
)

// Database instance
var db *sql.DB

// Create Fake DB for Quotes
var quotes = []*constant.Quotes{
	{ID: 1, Text: "There is only one thing that makes a dream impossible to achieve: the fear of failure.", Author: "Paulo Coelho", Category: []string{"Motivation"}},
	{ID: 2, Text: "Have no fear of perfection - you'll never reach it.", Author: "Salvador Dalí", Category: []string{"Motivation"}},
	{ID: 3, Text: "Whatever you are, be a good one.", Author: "Abraham Lincoln", Category: []string{"Motivation"}},
	{ID: 4, Text: "Make everything as simple as possible, but not simpler.", Author: "Albert Einstein", Category: []string{"Motivation"}},
	{ID: 4, Text: "Creativity is intelligence having fun.", Author: "Albert Einstein", Category: []string{"Motivation"}},
	{ID: 5, Text: "The best revenge is massive success.", Author: "Frank Sinatra", Category: []string{"Motivation"}},
	{ID: 6, Text: "If people never did silly things, nothing intelligent would ever get done.", Author: "Ludwig Wittgenstein", Category: []string{"Motivation"}},
	{ID: 7, Text: "The two most important days in your life are the day you are born and the day you find out why.", Author: "Mark Twain", Category: []string{"Motivation"}},
	{ID: 8, Text: "Pain is inevitable. Suffering is optional.", Author: "Haruki Murakami", Category: []string{"Motivation"}},
	{ID: 9, Text: "The only truly secure system is one that is powered off, cast in a block of concrete and sealed in a lead-lined room with armed guards.", Author: "Gene Spafford", Category: []string{"Cyber Security"}},
	{ID: 10, Text: "The more you know, the more you realize you know nothing.", Author: "Socrates", Category: []string{"Motivation"}},
	{ID: 11, Text: "If you reveal your secrets to the wind, you should not blame the wind for revealing them to the trees.", Author: "Kahlil Gibran", Category: []string{"Motivation"}},
	{ID: 12, Text: "The best way out is always through.", Author: "Robert Frost", Category: []string{"Motivation"}},
	{ID: 13, Text: "If you don't risk anything, you risk even more.", Author: "Erica Jong", Category: []string{"Motivation"}},
	{ID: 14, Text: "There are far, far better things ahead than any we leave behind.", Author: "C.S. Lewis", Category: []string{"Motivation"}},
	{ID: 15, Text: "Follow your inner moonlight. Don't hide the madness.", Author: "Allen Ginsberg", Category: []string{"Motivation"}},
	{ID: 16, Text: "The harder the conflict, the more glorious the triumph.", Author: "Thomas Paine", Category: []string{"Motivation"}},
	{ID: 17, Text: "Every strike brings me closer to the next home run.", Author: "Babe Ruth", Category: []string{"Motivation"}},
	{ID: 18, Text: "Everybody has talent, but ability takes hard work.", Author: "Michael Jordan", Category: []string{"Motivation"}},
}

// Create Fake DB for Authors
var authorGallery = []*constant.AuthorGallery{
	{ID: 1, Name: "Salvador Dali", Url: "salvador-dali.jpeg", Category: []string{"Artist"}},
}

// @Summary Get all Quotes
// @Description Get all Quotes
// @Tags Quotes
// @Accept json
// @Produce json
// @Success 200 {array} Quotes{}
// @Router /api/quotes [get]
func GetQuotes(ctx *fiber.Ctx) error {
	quotes := []models.Quotes{}
	database.DB.Db.Find(&quotes)

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
		if idErr := helper.CreateError("cannot parse id"); idErr != nil {
			return idErr
		}
	}

	quote := []models.Quotes{}
	database.DB.Db.Model(&quote).Where("id = ?", _id).Find(&quote)
	if len(quote) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "item is not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"item":    quote,
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

	var reQuotes []*constant.Quotes

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
		Text *string `json:"text"`
		// Author *string `json:"author"`
		// Category *[]string `json:"category"`
	}

	var body request

	if err := ctx.BodyParser(&body); err != nil {
		if bodyErr := helper.CreateError("body parse error"); bodyErr != nil {
			return bodyErr
		}
	}

	rows, err := db.Query("INSERT INTO quotes (text) VALUES($1)", body.Text)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("...")
	}
	defer rows.Close()
	result := request{}

	for rows.Next() {
		res := request{}
		if err := rows.Scan(&res.Text); err != nil {
			return err
		}

		result.Text = res.Text
	}

	fmt.Println(rows)

	// newQuote := &constant.Quotes{
	// 	ID:     uint(len(quotes) + 1),
	// 	Text:   *body.Text,
	// 	Author: *body.Author,
	// 	// Category: *body.Category,
	// }

	// quotes = append(quotes, newQuote)

	return ctx.Status(fiber.StatusOK).JSON("...")
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
		if err := helper.CreateError("params parse error"); err != nil {
			return err
		}
	}

	for i, t := range quotes {
		if t.ID == uint(_id) {
			quotes = append(quotes[0:i], quotes[i+1:]...)

			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"message": "Quote successfuly deleted",
			})
		}
	}

	return ctx.Status(fiber.StatusBadRequest).JSON("something went wrong")
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
		if idErr := helper.CreateError("cannot parse id"); idErr != nil {
			return idErr
		}
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

	var reAuthor []*constant.AuthorGallery

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
	author := &constant.AuthorGallery{
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
		if paramsErr := helper.CreateError("params parse error"); paramsErr != nil {
			return paramsErr
		}
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
