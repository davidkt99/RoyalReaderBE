package db

import (
	"github.com/davidkt99/RoyalReaderBE/models"
)

func InsertBook(book models.Book) int64 {
	insertStmt := `insert into "books"("book_name", "book_url") values($1, $2) RETURNING book_id`
	var id int64
	e := db.QueryRow(insertStmt, book.Name, book.Url).Scan(&id)
	CheckError(e)
	return id
}
