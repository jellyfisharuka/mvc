package main

import (
	"mvc/controllers"
	"mvc/initializer"
	"os"
    "github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializer.LoadEnvariables()
	initializer.ConnectionDatabase()
	initializer.SyncDB()
}
func main() {
    //load  templates
	engine:=html.New("./views",".tmpl")
	//setup app
	app := fiber.New(fiber.Config{
		Views:engine,
	})
	//configure app
	app.Static("/", "./public")
    //routes
	app.Get("/", controllers.POstsIndex)
    //start app
	app.Listen(":"+os.Getenv("PORT"))
}
