package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func GetUserDatails(ctx *fiber.Ctx) error {
	userDetails := User{
		Name:   "Sabita Neupane",
		Age:    "26",
		Gender: "Female",
	}
	return ctx.Status(fiber.StatusOK).JSON(userDetails)
}

func CreateUserDatails(ctx *fiber.Ctx) error {
	body := new(User)
	err := ctx.BodyParser(body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	userDetails := User{
		Name:   body.Name,
		Age:    body.Age,
		Gender: body.Gender,
	}

	details := map[string]interface{}{
		"message": "User Details Created Successfully",
		"data":    userDetails,
	}

	return ctx.Status(fiber.StatusOK).JSON(details)
}

func ServiceStatus(ctx *fiber.Ctx) error {
	msg := map[string]interface{}{
		"message": "All good from Fibergo Learn app.",
	}
	return ctx.JSON(msg)
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	app.Get("/service-status", ServiceStatus)

	// grouping same/similar routes
	fibergoApp := app.Group("/user")
	fibergoApp.Get("/", GetUserDatails)
	fibergoApp.Post("/", CreateUserDatails)

	app.Listen(":80")
}
