package db

func InsertWeeklyChapterUpdate(bookId int64, chapterId int64, dayOfWeek int64) {
	insertStmt := `insert into "weekly_updates"("book_key", "chapter_key", "day_of_week") values($1, $2, $3)`

	_, e := db.Exec(insertStmt, bookId, chapterId, dayOfWeek)
	CheckError(e)
}
