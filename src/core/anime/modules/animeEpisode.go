package usecases

import (
	"api/src/common/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func AnimeEpisode() string {
	cfg := config.NewConfig()
	return cfg.Anime.Hls
}

func AnimeEpisodeHandler(c *fiber.Ctx) error {
	baseUrl := AnimeEpisode()
	name := c.Params("name")
	episode := c.Params("ep")

	animeDir := filepath.Join("src", "assets", name)
	outputFile := filepath.Join(animeDir, fmt.Sprintf("%s.m3u8", episode))

	if fileExists(outputFile) {
		c.Set("Access-Control-Allow-Origin", "*") 
		c.Set("Content-Type", "image/webp") 
		c.Set("Accept", "*/*")
		c.Set("Referer", "https://www.anroll.net/")
		return c.SendFile(outputFile)
	}

	url := fmt.Sprintf("%s/animes/%s/%s.mp4/media-1/stream.m3u8", baseUrl, name, episode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request: " + err.Error())
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Referer", "https://www.anroll.net/")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error making GET request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(resp.StatusCode).SendString("Error: received status " + resp.Status)
	}

	if err := os.MkdirAll(animeDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating directory: " + err.Error())
	}

	if err := saveFile(resp.Body, outputFile); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving response body: " + err.Error())
	}

	c.Set("Access-Control-Allow-Origin", "*") 
	c.Set("Content-Type", "image/webp")
	return c.SendFile(outputFile)
}

func saveFile(src io.Reader, dest string) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
