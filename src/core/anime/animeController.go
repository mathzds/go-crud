package anime

import (
	usecases "api/src/core/anime/modules"

	"github.com/gofiber/fiber/v2"
)

func AnimeController(app *fiber.App) {
	app.Get("/anime/search/", usecases.AnimeSearchHandler)
	app.Get("/anime/releases", usecases.AnimeReleasesHandler)
	app.Get("/animes/info/:id/:page/:order", usecases.AnimeInfoHandlerEpisodes)
	app.Get("/animes/info/data", usecases.AnimeDataHandler)
}
