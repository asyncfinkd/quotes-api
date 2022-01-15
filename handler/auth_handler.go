package handler

import (
	"quotes-api/config"
	"quotes-api/database"
	models "quotes-api/models/quotes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY = []byte("gosecretkey")

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Signin(ctx *fiber.Ctx) error {
	collection := database.Global().Db.Collection("users")
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	quote := new(request)
	var result models.Users

	if err := ctx.BodyParser(quote); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	filter := bson.D{{"email", quote.Email}}

	if err := collection.FindOne(ctx.Context(), filter).Decode(&result); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Credentials incorrect.",
		})
	}

	if !CheckPasswordHash(quote.Password, result.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Invalid password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET_KEY")))

	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":      true,
		"message":      "you logged successfully.",
		"access_token": t,
	})
}
