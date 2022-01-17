package handler

import (
	"fmt"
	"log"
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

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
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
	claims["iat"] = 0
	claims["_id"] = result.ID
	claims["name"] = result.Name
	claims["email"] = result.Email
	claims["logged"] = true

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

func Signup(ctx *fiber.Ctx) error {
	collection := database.Global().Db.Collection("users")
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	user := new(request)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).SendString(err.Error())
	}

	user.Password = getHash([]byte(user.Password))
	fmt.Println(user.Password)

	xml, err := collection.InsertOne(ctx.Context(), user)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	fmt.Println(xml)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "user successfuly added.",
	})
}
