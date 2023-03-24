package main

import (
	"github.com/davidkt99/RoyalReaderBE/db"
	_ "github.com/lib/pq"
)

func main() {
	//*	Database Start
	db.DBSetup()
	defer db.DBShutDown()

	//*		given url for book scrape all chapters and insert into database
	// bookTitle, bookUrl, allChapters := scraper.ScrapeBook("https://www.royalroad.com/fiction/65145/elysiums-multiverse")
	// book := models.Book{Name: bookTitle, Url: bookUrl}
	// bookId := db.InsertBook(book)
	// chapterIdList := db.InsertAllChapters(allChapters, bookId)

	// //* CronJobs Start
	// jobs.CronJobsSetup()

	// //*	Services Start
	// services.StartServices()

	db.QueryChapterById(278)

}
