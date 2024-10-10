package usecases

import (
	"api/src/common/config"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func AnimeSearch(query string) string {
	cfg := config.NewConfig()
	anime := cfg.Anime
	return anime.SearchURL(query)
}

func AnimeSearchHandler(c *fiber.Ctx) (err error) {
	query := c.Query("name")
	url := AnimeSearch(query)
	agent := fiber.Get(url)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	var something fiber.Map
	err = json.Unmarshal(body, &something)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	return c.Status(statusCode).JSON(something)
}
