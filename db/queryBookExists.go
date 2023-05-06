package db

func QueryBookExists(id int64) bool {
	insertStmt := `
	select book_id
	from books
	where books.book_id=$1`

	rows, e := db.Query(insertStmt, id)
	CheckError(e)

	defer rows.Close()
	if rows.Next() {
		return true
	} else {
		return false
	}
}
