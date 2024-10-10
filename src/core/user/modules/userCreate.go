package usecases

import (
	handlers "api/src/common/handlers"
	models "api/src/common/types"

	"github.com/gofiber/fiber/v2"
)

func CreateUser (user models.User) (models.User, error) {
	result := handlers.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func CreateUserHandler(c *fiber.Ctx) error {
	var newUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid data",
		})
	}
	createdUser, err := CreateUser(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "cannot create user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(createdUser)
}