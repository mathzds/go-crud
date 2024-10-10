package anime

import (
	usecases "api/src/core/anime/modules"

	"github.com/gofiber/fiber/v2"
)

func AnimeController(app *fiber.App) {
	app.Get("/anime/search/", usecases.AnimeSearchHandler)
	app.Get("/anime/releases", usecases.AnimeReleasesHandler)
}
