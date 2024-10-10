package usecases

import (
	"api/src/common/handlers"
	models "api/src/common/types"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UpdateUser (id uint, updateData models.User) (models.User, error) {
	var user models.User
	if err := handlers.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	
	if err := handlers.DB.Model(&user).Updates(updateData).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUserHandler (c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid id",
		})
	}

	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid data",
		})
	}

	user, err := UpdateUser(uint(id), updatedUser)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error" : "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "cannot update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)

}