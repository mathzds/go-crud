package routes

import (
	user "api/src/core/user"
	anime "api/src/core/anime"	
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) error {
	user.UserController(app)
	anime.AnimeController(app)
	return nil
}