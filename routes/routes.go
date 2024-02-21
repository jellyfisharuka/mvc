package routes

import (
	"mvc/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) error {
	app.Get("/", controllers.POstsIndex)
	return nil
}