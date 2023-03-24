package main

import (
	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/scraper"
	_ "github.com/lib/pq"
)

func main() {
	db.DBSetup()
	defer db.DBShutDown()

	// bookTitle, bookUrl, allChapters := scraper.ScrapeBook("https://www.royalroad.com/fiction/65145/elysiums-multiverse")
	// book := models.Book{Name: bookTitle, Url: bookUrl}
	// bookId := db.InsertBook(book)
	// chapterIdList := db.InsertAllChapters(allChapters, bookId)

	// fmt.Println(bookId)
	// fmt.Println(chapterIdList)

	scraper.CheckNewChapters("https://www.royalroad.com/fiction/65145/elysiums-multiverse")
	// services.StartServices()

}
