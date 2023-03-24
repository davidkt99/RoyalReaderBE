package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
)

func QueryAllBooks() []models.Book {
	insertStmt := `
	select book_id, book_name, book_url
	from books`

	var books []models.Book

	rows, e := db.Query(insertStmt)
	CheckError(e)

	defer rows.Close()
	for rows.Next() {
		book := models.Book{}
		e := rows.Scan(&book.Id, &book.Name, &book.Url)
		CheckError(e)

		books = append(books, book)
	}

	//* Rows returned empty
	if len(books) < 1 {
		fmt.Println("No Books Found")
	} else {
		fmt.Println("Queried # of Books: ", len(books))
	}

	return books
}
