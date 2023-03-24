package db

import "fmt"

type ChapterNameAndId struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func QueryChapterNameAndIdsByBookId(bookId int64) []ChapterNameAndId {
	insertStmt := `
	select chapter_id, chapter_name
	from books
	join chapters on books.book_id=chapters.book_key
	where books.book_id=$1`

	var chapters []ChapterNameAndId

	rows, e := db.Query(insertStmt, bookId)
	CheckError(e)

	defer rows.Close()
	for rows.Next() {
		chapter := ChapterNameAndId{}
		e := rows.Scan(&chapter.Id, &chapter.Name)
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
