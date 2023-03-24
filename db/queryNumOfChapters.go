package db

import (
	"fmt"
)

func QueryNumOfChapters(bookId int64) int {
	insertStmt := `
	select count(*)
	from books
	join chapters on books.book_id=chapters.book_key
	where books.book_id=$1`

	var count int

	rows, e := db.Query(insertStmt, bookId)
	CheckError(e)

	defer rows.Close()
	rows.Next()
	eScan := rows.Scan(&count)
	CheckError(eScan)

	fmt.Println("Num of Chapters: ", count)

	return count
}
