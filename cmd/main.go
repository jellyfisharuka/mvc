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
// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
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
// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
	   "data": "Server is up and running",
	}
 
	if err := c.JSON(res); err != nil {
	   return err
	}
 
	return nil
 }