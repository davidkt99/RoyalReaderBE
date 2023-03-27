package jobs

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/scraper"
)

func NewChaptersJob() {
	books := db.QueryAllBooks()
	for _, book := range books {
		currChapterNum := scraper.ScrapeChatperNum(book.Url)
		storedChapterNum := db.QueryNumOfChapters(book.Id)
		numOfNewChapters := currChapterNum - storedChapterNum
		if numOfNewChapters > 0 {
			fmt.Println("There are new Chapters for ", book.Name, book.Id, book.Url, " - Num of New Chapters: ", numOfNewChapters)
			chapters := scraper.ScrapeNewChapters(book.Url, storedChapterNum)
			for _, chapter := range chapters {
				chapter.BookId = book.Id
				db.InsertChapter(chapter)
			}
		} else {
			fmt.Println("NO new Chapters for ", book.Name, book.Id, book.Url)

		}
	}
}
