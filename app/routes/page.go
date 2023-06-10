package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	short := &Response{}
	return c.Render("index", fiber.Map{
		"url": short.CustomShort,
	})
}