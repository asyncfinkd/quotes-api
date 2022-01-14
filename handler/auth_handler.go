// package handler

// import (
// 	"quotes-api/config"
// 	// "quotes-api/constant"
// 	"quotes-api/helper"
// 	"strings"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt"
// 	"golang.org/x/crypto/bcrypt"
// )

// // Create Fake DB for Users
// // var users = []*constant.User{
// // 	{ID: 0, Name: "Nika Shamiladze", Email: "zxc@gmail.com", Password: "$2a$12$0IelvstJ1QLvFZOH8GM8dOuzu/ouhBNE2DJ3GpfK79dzZ4mO5JuHu", Role: "user"},
// // }

// // Here we use (bcrypt)
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func Signin(ctx *fiber.Ctx) error {
// 	type Request struct {
// 		Email    *string `json:"email"`
// 		Password *string `json:"password"`
// 	}

// 	var body Request

// 	if err := ctx.BodyParser(&body); err != nil {
// 		if bodyErr := helper.CreateError("cannot parse body"); bodyErr != nil {
// 			return bodyErr
// 		}
// 	}

// 	if !strings.Contains(*body.Email, "@") {
// 		if err := helper.CreateError("email address you entered doesn't contain '@' sign"); err != nil {
// 			return err
// 		}
// 	}

// 	if len(*body.Password) < 6 {
// 		if err := helper.CreateError("password you entered is too short"); err != nil {
// 			return err
// 		}
// 	}

// 	for _, t := range users {
// 		if !CheckPasswordHash(*body.Password, t.Password) {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Invalid password", "access_token": nil})
// 		}

// 		loginValidate := t.Email == *body.Email
// 		if loginValidate {
// 			token := jwt.New(jwt.SigningMethodHS256)
// 			claims := token.Claims.(jwt.MapClaims)
// 			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// 			t, err := token.SignedString([]byte(config.Config("SECRET_KEY")))

// 			if err != nil {
// 				return ctx.SendStatus(fiber.StatusInternalServerError)
// 			}

// 			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
// 				"success":      true,
// 				"message":      "Congratulation, you logged succesfully",
// 				"access_token": t,
// 			})
// 		}
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON("something is wrong")
// }
package handler
