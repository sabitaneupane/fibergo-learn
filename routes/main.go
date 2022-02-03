package routes

import (
	"fibergo-learn/controllers"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/", controllers.GetHelloWorld)
	app.Get("/service-status", controllers.ServiceStatus)
	app.Get("/api", controllers.GetStaticPage)

	// grouping same/similar routes
	fibergoApp := app.Group("/user")
	fibergoApp.Get("/", controllers.GetUserDetails)
	fibergoApp.Post("/", controllers.CreateUserDetails)
}
