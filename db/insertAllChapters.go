package db

import "github.com/davidkt99/RoyalReaderBE/models"

func InsertAllChapters(chapters []models.Chapter, bookId int64) {

	for _, chapter := range chapters {
		chapter.BookId = bookId
		InsertChapter(chapter)
	}
}
