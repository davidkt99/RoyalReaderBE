package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
)

func QueryChapterById(chapterId int64) models.Chapter {
	insertStmt := `
	select chapter_id, chapter_name, chapter_content, book_key
	from chapters
	where chapters.chapter_id=$1`

	var chapter models.Chapter

	rows, e := db.Query(insertStmt, chapterId)
	CheckError(e)

	defer rows.Close()
	rows.Next()
	eScan := rows.Scan(&chapter.Id, &chapter.Name, &chapter.Content, &chapter.BookId)
	CheckError(eScan)

	fmt.Println("Chapter Found: ", chapter.Id)

	return chapter
}
