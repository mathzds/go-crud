package user

import (
	usecases "api/src/core/user/modules"

	"github.com/gofiber/fiber/v2"
)

func UserController(app *fiber.App) {
	app.Post("/user", usecases.CreateUserHandler)
	app.Delete("/user/:id", usecases.DeleteUserHandler)
	app.Put("/user/:id", usecases.UpdateUserHandler)
	app.Get("/user/:id", usecases.FindUserHandler)
}
