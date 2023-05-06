package db

import (
	"github.com/davidkt99/RoyalReaderBE/models"
)

func InsertBook(book models.Book) int64 {
	insertStmt := `insert into "books"("book_id", "book_name", "author", "image_url") values($1, $2, $3, $4) RETURNING book_id`
	var id int64
	e := db.QueryRow(insertStmt, book.Id, book.Name, book.Author, book.ImageUrl).Scan(&id)
	CheckError(e)
	return id
}
