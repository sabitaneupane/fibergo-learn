package routes

import (
	"fibergo-learn/controllers"
	"fibergo-learn/routes/userroutes"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", controllers.GetStaticPage)
	app.Get("/hello", controllers.GetHelloWorld)
	app.Get("/service-status", controllers.ServiceStatus)

	userroutes.Router(app)
}
