package usecases

import (
	"api/src/common/config"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
)
func replaceParams(template string, params map[string]string) string {
	result := template
	for key, value := range params {
		result = strings.ReplaceAll(result, "${"+key+"}", value)
	}
	return result
}

func AnimeInfoEpisodes() string {
	cfg := config.NewConfig()
	anime := cfg.Anime
	return anime.AnimeEpisode
}

func AnimeInfoHandlerEpisodes(c *fiber.Ctx) (err error) {
	baseUrl := AnimeInfoEpisodes()

	id := c.Params("id")
	page := c.Params("page")
	order := c.Params("order")

	url := replaceParams(baseUrl, map[string]string{
		"id":    id,
		"page":  page,
		"order": order,
	})

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