package main

import (
	"github.com/davidkt99/RoyalReaderBE/scraper"
	_ "github.com/lib/pq"
)

func main() {
	// scraper.ScrapeChapter("https://www.royalroad.com/fiction/65359/ashes-of-heaven-book-one/chapter/1138407/chapter-24-the-great-dao-of-scamming")

	scraper.ScrapeBook("https://www.royalroad.com/fiction/65324/wolf-of-the-blood-moon")

	// db.DBSetup()

	// services.StartServices()

}
