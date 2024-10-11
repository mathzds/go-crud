package config

import (
	"net/url"
	"strconv"
	"strings"
)

type Config struct {
	Port  string
	Anime AnimeConfig
}

type AnimeConfig struct {
	ImagesThumbnail string
	ImagesEpisode   string
	Search          string
	Releases        string
	Anime           string
	AnimeEpisode    string
	Hls            string
}

func NewConfig() *Config {
	return &Config{
		Port: "3000",
		Anime: AnimeConfig{
			Hls:            "https://cdn-zenitsu-2-gamabunta.b-cdn.net/cf/hls",
			ImagesThumbnail: "https://static.anroll.net/images/${type}/capas/${slug}.jpg",
			ImagesEpisode:   "https://static.anroll.net/images/animes/screens/${slug}/${number}.jpg",
			Search:          "https://api-search.anroll.net/data?q=${query}",
			Releases:        "https://www.anroll.net/",
			Anime:           "https://www.anroll.net/${generic}",
			AnimeEpisode:    "https://apiv3-prd.anroll.net/animes/${id}/episodes?page=${page}&order=${order}",
		},
	}
}
func (a AnimeConfig) ImagesThumbnailURL(typeParam, slug string) string {
	return replaceParams(a.ImagesThumbnail, map[string]string{"type": typeParam, "slug": slug})
}


func (a AnimeConfig) ImagesEpisodeURL(slug string, number int) string {
	return replaceParams(a.ImagesEpisode, map[string]string{
		"slug":   slug,
		"number": strconv.Itoa(number),
	})
}
func (a AnimeConfig) SearchURL(query string) string {
	return replaceParams(a.Search, map[string]string{"query": url.QueryEscape(query)})
}

func replaceParams(template string, params map[string]string) string {
	result := template
	for key, value := range params {
		result = strings.ReplaceAll(result, "${"+key+"}", value)
	}
	return result
}
