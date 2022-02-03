package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func ServiceStatus(ctx *fiber.Ctx) error {
	msg := map[string]interface{}{
		"message": "All good from Fibergo Learn app.",
	}
	return ctx.JSON(msg)
}
