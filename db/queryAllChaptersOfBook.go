package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
)

func QueryAllChaptersOfBook(bookId int64) []models.Chapter {
	insertStmt := `
	select chapter_id, chapter_name, chapter_content, read
	from books
	join chapters on books.book_id=chapters.book_key
	where books.book_id=$1`

	var chapters []models.Chapter

	rows, e := db.Query(insertStmt, bookId)
	CheckError(e)

	defer rows.Close()
	for rows.Next() {
		chapter := models.Chapter{}
		chapter.BookId = bookId
		e := rows.Scan(&chapter.Id, &chapter.Name, &chapter.Content, &chapter.Read)
		CheckError(e)

		chapters = append(chapters, chapter)
	}

	//* Rows returned empty
	if len(chapters) < 1 {
		fmt.Println("Book not found: ", bookId)
	} else {
		fmt.Println("Queried # of chapters: ", len(chapters))
	}

	return chapters
}
