package db

import "github.com/davidkt99/RoyalReaderBE/models"

func InsertChapter(chapter models.Chapter) int64 {
	insertStmt := `insert into "chapters"("chapter_name", "chapter_content", "read", "book_key") values($1, $2, false, $3) RETURNING chapter_id`
	var id int64
	e := db.QueryRow(insertStmt, chapter.Name, chapter.Content, chapter.BookId).Scan(&id)
	CheckError(e)
	return id
}
