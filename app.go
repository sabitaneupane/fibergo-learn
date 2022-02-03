package main

import (
	"fibergo-learn/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.Router(app)

	app.Listen(":80")
}
