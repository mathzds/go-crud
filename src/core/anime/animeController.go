package anime

import (
	usecases "api/src/core/anime/modules"

	"github.com/gofiber/fiber/v2"
)

func AnimeController(app *fiber.App) {
	app.Get("/animes/search/", usecases.AnimeSearchHandler)
	app.Get("/animes/releases", usecases.AnimeReleasesHandler)
	app.Get("/animes/info/:id/:page/:order", usecases.AnimeInfoHandlerEpisodes)
	app.Get("/animes/info/data", usecases.AnimeDataHandler)
	app.Get("animes/stream/:name/:ep", usecases.AnimeEpisodeHandler)
}
