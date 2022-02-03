package controllers

import (
	"fibergo-learn/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUserDetails(ctx *fiber.Ctx) error {
	body := new(model.User)
	err := ctx.BodyParser(body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	userDetails := model.User{
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
	userDetails := model.User{
		Name:   "Sabita Neupane",
		Age:    "26",
		Gender: "Female",
	}
	return ctx.Status(fiber.StatusOK).JSON(userDetails)
}
