package userroutes

import (
	"fibergo-learn/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/user", controllers.GetUserDetails)
	app.Post("/user", controllers.CreateUserDetails)
}
