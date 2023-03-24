package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
)

func QueryBookById(bookId int64) models.Book {
	insertStmt := `
	select book_id, book_name, book_url
	from books
	where books.book_id=$1`

	var book models.Book

	rows, e := db.Query(insertStmt, bookId)
	CheckError(e)

	defer rows.Close()
	rows.Next()
	eScan := rows.Scan(&book.Id, &book.Name, &book.Url)
	CheckError(eScan)

	fmt.Println("Book Found: ", book.Id)

	return book
}
