package usecases

import (
	"api/src/common/config"
	"api/src/common/utils"

	"github.com/gofiber/fiber/v2"
)

func AnimeReleases() string {
	cfg := config.NewConfig()
	anime := cfg.Anime
	return anime.Releases
}

func AnimeReleasesHandler(c *fiber.Ctx) error {
	url := AnimeReleases()
	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	releases, err := utils.ExtractData(string(body), "release")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(releases)
}