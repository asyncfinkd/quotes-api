package handler

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var users = []*User{
	{ID: 0, Name: "Nika Shamiladze", Email: "zxc@gmail.com", Password: "123123123", Role: "user"},
}

func Signin(ctx *fiber.Ctx) error {
	type Request struct {
		Email    *string `json:"email"`
		Password *string `json:"password"`
	}

	var body Request

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
	}

	if !strings.Contains(*body.Email, "@") {
		return ctx.Status(fiber.StatusBadRequest).JSON("email address you entered doesn't contain '@' sign")
	}

	if len(*body.Password) < 6 {
		return ctx.Status(fiber.StatusBadRequest).JSON("password you entered is too short")
	}

	for _, t := range users {
		loginValidate := t.Email == *body.Email && t.Password == *body.Password
		if loginValidate {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			t, err := token.SignedString([]byte("secret"))

			if err != nil {
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}

			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"success":      true,
				"message":      "Congratulation, you logged succesfully",
				"access_token": t,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON("something is wrong")
}
