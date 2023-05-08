package db

import (
	"time"

	"github.com/davidkt99/RoyalReaderBE/models"
)

func InsertChapter(chapter models.Chapter) int64 {
	var id int64
	insertStmt := `insert into "chapters"("chapter_name", "chapter_content", "book_key", "date_added") values($1, $2, $3, $4) RETURNING chapter_id`

	e := db.QueryRow(insertStmt, chapter.Name, chapter.Content, chapter.BookId, time.Now().Format("2006-01-02 15:04:05")).Scan(&id)
	CheckError(e)

	return id
}
