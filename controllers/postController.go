package controllers

import "github.com/gofiber/fiber/v2"

func POstsIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, Worl!")
}
