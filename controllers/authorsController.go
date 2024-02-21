package controllers

import "github.com/gofiber/fiber/v2"

func AuthorIndex(c *fiber.Ctx) error {
	return c.Render("authors/index", fiber.Map{})
}