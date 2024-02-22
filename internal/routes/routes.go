package routes

import (
	"mvc/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) error {
	app.Get("/", controllers.POstsIndex)
	app.Get("/", controllers.AuthorIndex)

	return nil
}
