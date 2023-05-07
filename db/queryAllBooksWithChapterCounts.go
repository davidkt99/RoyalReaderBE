package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
)

func QueryAllBooksWithChapterCounts() []models.BookWithChapterCount {
	insertStmt := `
	select book_id, book_name, author, image_url, (select count('x') from chapters where chapters.book_key=books.book_id)
from books`

	var books []models.BookWithChapterCount

	rows, e := db.Query(insertStmt)
	CheckError(e)

	defer rows.Close()
	for rows.Next() {
		book := models.BookWithChapterCount{}
		e := rows.Scan(&book.Id, &book.Name, &book.Author, &book.ImageUrl, &book.NumOfChapters)
		CheckError(e)

		books = append(books, book)
	}

	//* Rows returned empty
	if len(books) < 1 {
		fmt.Println("No Books Found")
		return []models.BookWithChapterCount{}
	} else {
		fmt.Println("Queried # of Books: ", len(books))
	}

	return books
}
