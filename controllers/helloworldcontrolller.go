package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHelloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello World")
}
