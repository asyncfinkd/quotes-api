package handler

import "github.com/gofiber/fiber/v2"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var users = []*User{
	{ID: 0, Name: "Nika Shamiladze", Email: "zxc@gmail.com", Password: "123", Role: "user"},
}

func Signin(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("...")
}
