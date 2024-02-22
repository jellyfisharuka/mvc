package main

import (
	"fmt"
	"mvc/internal/initializer"
	"mvc/internal/routes"
	"os"

	"github.com/gofiber/contrib/swagger"
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
//swagger
cfg := swagger.Config{
    BasePath: "/",
    FilePath: "./docs/swagger.json",
    Path:     "swagger",
    Title:    "Swagger API Docs",
}
app.Use(swagger.New(cfg))
}
