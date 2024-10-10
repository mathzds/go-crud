package usecases

import (
	"api/src/common/config"
	models "api/src/common/types"
	"encoding/json"
	"fmt"
	"strings"

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

	releases, err := ExtractReleaseData(string(body))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.Status(statusCode).JSON(releases)
}

func ExtractReleaseData(html string) ([]models.Release, error) {
	var jsonString string
	if splitData := strings.Split(html, "__NEXT_DATA__"); len(splitData) > 1 {
		if jsonPart := strings.Split(splitData[1], `type="application/json">`); len(jsonPart) > 1 {
			jsonString = strings.Split(jsonPart[1], "</script>")[0]
		} else {
			return nil, fmt.Errorf("__NEXT_DATA__")
		}
	} else {
		return nil, fmt.Errorf("__NEXT_DATA__")
	}

	var data models.Data
	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		return nil, err
	}

	return data.Props.PageProps.Data.DataReleases, nil
}
