package usecases

import (
	"api/src/common/config"
	"api/src/common/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AnimeData () string {
	cfg := config.NewConfig()
	anime := cfg.Anime.Anime
	return anime
}
func AnimeDataHandler(c *fiber.Ctx) error {
	generic := c.Query("g")
	if generic == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "generic is required",
		})
	}


	anime := AnimeData()
	url := strings.ReplaceAll(anime, "${generic}", generic)
	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	extractedAnime, err := utils.ExtractData(string(body), "anime")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(extractedAnime)
}