package usecases

import (
	"api/src/common/handlers"
	models "api/src/common/types"

	"github.com/gofiber/fiber/v2"
)

func FindUser(id uint) (models.User, error) {
	var user models.User
	if err := handlers.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func FindUserHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid id",
		})
	}

	user, err := FindUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "user not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}