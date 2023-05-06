package jobs

import (
	"fmt"
	"strconv"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/scraper"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func NewChaptersJob() {
	books := db.QueryAllBooks()
	for _, book := range books {
		checkForNewChapters(book)
	}
}

func checkForNewChapters(book models.Book) {
	url := util.ROYAL_ROAD_URL + strconv.FormatInt(book.Id, 10)

	currChapterNum := scraper.ScrapeChatperNum(url)
	storedChapterNum := db.QueryNumOfChapters(book.Id)
	numOfNewChapters := currChapterNum - storedChapterNum

	if numOfNewChapters > 0 {
		getNewChapters(book, url, storedChapterNum, numOfNewChapters)
		fmt.Println("There are new Chapters for ", book.Name, book.Id, " - Num of New Chapters: ", numOfNewChapters)
	} else {
		fmt.Println("NO new Chapters for ", book.Name, book.Id)

	}
}

func getNewChapters(book models.Book, url string, storedChapterNum int, numOfNewChapters int) {
	chapters := scraper.ScrapeNewChapters(url, storedChapterNum)

	for _, chapter := range chapters {
		chapter.BookId = book.Id
		db.InsertChapter(chapter)
	}
}
