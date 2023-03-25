package db

import "github.com/davidkt99/RoyalReaderBE/models"

func InsertChapter(chapter models.Chapter) {
	insertStmt := `insert into "chapters"("chapter_name", "chapter_content", "book_key") values($1, $2, $3) RETURNING chapter_id`

	_, e := db.Exec(insertStmt, chapter.Name, chapter.Content, chapter.BookId)
	CheckError(e)
}
