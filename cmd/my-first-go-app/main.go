package main

import (
	"fmt"

	"github.com/Shomey9/myFirstGoApp/config"
	"github.com/Shomey9/myFirstGoApp/internal/scraper"
	"github.com/Shomey9/myFirstGoApp/internal/utils"
)

func main() {
	logger := utils.InitLogger()
	logger.Info("Starting RSS feed scraper")

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("error loading config: %v", err)
		return
	}

	fmt.Println("Starting RSS scraping...")
	scraper.ScrapeFeeds(cfg.Feeds)
	fmt.Println("Scrapping complete!")

}
