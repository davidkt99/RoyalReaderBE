package jobs

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/scraper"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func NewChaptersJob() {
	books := db.QueryAllBooks()
	for _, book := range books {
		url := util.ROYAL_ROAD_URL + string(book.Id)
		currChapterNum := scraper.ScrapeChatperNum(url)
		storedChapterNum := db.QueryNumOfChapters(book.Id)
		numOfNewChapters := currChapterNum - storedChapterNum
		if numOfNewChapters > 0 {
			fmt.Println("There are new Chapters for ", book.Name, book.Id, " - Num of New Chapters: ", numOfNewChapters)
			chapters := scraper.ScrapeNewChapters(url, storedChapterNum)
			for _, chapter := range chapters {
				chapter.BookId = book.Id
				db.InsertChapter(chapter)
			}
		} else {
			fmt.Println("NO new Chapters for ", book.Name, book.Id)

		}
	}
}
