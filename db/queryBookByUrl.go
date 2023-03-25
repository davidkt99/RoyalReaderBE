package db

func QueryBookByUrl(url string) bool {
	insertStmt := `
	select book_id, book_name, book_url
	from books
	where books.book_url=$1`

	rows, e := db.Query(insertStmt, url)
	CheckError(e)

	defer rows.Close()
	if rows.Next() {
		return true
	} else {
		return false
	}
}
