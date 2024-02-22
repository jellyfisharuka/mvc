package controllers

import "github.com/gofiber/fiber/v2"

func POstsIndex(c *fiber.Ctx) error {
	return c.Render("posts/index", fiber.Map{})
}
