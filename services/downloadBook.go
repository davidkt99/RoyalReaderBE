package services

import (
	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/scraper"
	"github.com/davidkt99/RoyalReaderBE/util"
)

// *		given url for book scrape all chapters and insert into database
// TODO: Add error checking
func DownloadBook(url string) string {

	//* Checking if Book already exists
	requestedBookId := util.BookIdFromUrl(url)
	if db.QueryBookExists(requestedBookId) {
		return "We already have this book in our catalogue"
	}

	book, allChapters := scraper.ScrapeBook(url)
	bookId := db.InsertBook(book)
	success := db.InsertAllChapters(allChapters, bookId)
	if success {
		return "This book was successfully added to the catalogue"
	} else {
		return "There was an issue adding this book, Please try again later"
	}
}
