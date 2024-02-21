package main

import (
	"fmt"
	"mvc/initializer"
	"os"
    "mvc/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
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
	routes.Routes(app)
    //start app
	err := app.Listen(":" + os.Getenv("PORT"))
if err != nil {
    fmt.Println("Error starting server:", err)
}
}
