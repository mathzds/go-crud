package usecases

import (
	"api/src/common/handlers"
	models "api/src/common/types"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteUser ( id uint) error {
	result := handlers.DB.Unscoped().Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUserHandler (c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid id",
		})
	}

	if err := DeleteUser(uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error" : "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "cannot delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "user deleted successfully",
	})
}