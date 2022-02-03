package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type User struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func getUserDetails(ctx *fiber.Ctx) error {
	userDetails := User{
		Name:   "Sabita Neupane",
		Age:    "26",
		Gender: "Female",
	}
	return ctx.Status(fiber.StatusOK).JSON(userDetails)
}

func createUserDetails(ctx *fiber.Ctx) error {
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

func serviceStatus(ctx *fiber.Ctx) error {
	msg := map[string]interface{}{
		"message": "All good from Fibergo Learn app.",
	}
	return ctx.JSON(msg)
}

func helloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World")
}

func getStaticPage(ctx *fiber.Ctx) error {
	return ctx.Render("../../public/index.html", fiber.Map{})
}

func main() {
	app := fiber.New()
	app.Use(logger.New()) // Middleware to log API request in terminal

	app.Get("/", helloWorld)
	app.Get("/service-status", serviceStatus)
	app.Get("/api", getStaticPage)

	// grouping same/similar routes
	fibergoApp := app.Group("/user")
	fibergoApp.Get("/", getUserDetails)
	fibergoApp.Post("/", createUserDetails)

	app.Listen(":80")
}
