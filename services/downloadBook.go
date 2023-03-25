package services

import (
	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/scraper"
)

// *		given url for book scrape all chapters and insert into database
// TODO: Add error checking
func DownloadBook(url string) bool {

	bookTitle, bookUrl, allChapters := scraper.ScrapeBook(url)
	book := models.Book{Name: bookTitle, Url: bookUrl}
	bookId := db.InsertBook(book)
	success := db.InsertAllChapters(allChapters, bookId)
	return success
}
