package main

import (
	"api/src/common/handlers"
	routes "api/src/common/routes"

	"github.com/gofiber/fiber/v2"
)

func main () {
	app := fiber.New()
	handlers.InitializeDatabase()
	routes.InitRoutes(app)
	app.Listen(":3000")
}