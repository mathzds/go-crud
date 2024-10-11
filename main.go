package main

import (
	"api/src/common/handlers"
	routes "api/src/common/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main () {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Acess-Control-Allow-Origin",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	handlers.InitializeDatabase()
	routes.InitRoutes(app)
	app.Listen(":3000")
}