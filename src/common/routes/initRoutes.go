package routes

import (
	routes "api/src/core/user"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) error {
	routes.UserController(app)
	return nil
}