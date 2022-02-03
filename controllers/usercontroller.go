package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func CreateUserDetails(ctx *fiber.Ctx) error {
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

func GetUserDetails(ctx *fiber.Ctx) error {
	userDetails := User{
		Name:   "Sabita Neupane",
		Age:    "26",
		Gender: "Female",
	}
	return ctx.Status(fiber.StatusOK).JSON(userDetails)
}
