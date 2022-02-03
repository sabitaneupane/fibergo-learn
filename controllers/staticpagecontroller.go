package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetStaticPage(ctx *fiber.Ctx) error {
	return ctx.Render("./public/index.html", fiber.Map{})
}
