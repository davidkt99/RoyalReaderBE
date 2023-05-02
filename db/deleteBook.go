package db

import "database/sql"

func DeleteBook(id int64) sql.Result {
	_, e := DeleteChapters(id)
	if e != nil {
		println(e.Error())
	}

	insertStmt := `
	DELETE FROM books where book_id=$1;
	`
	result, e := db.Exec(insertStmt, id)
	if e != nil {
		println(e.Error())
	}

	return result
}

func DeleteChapters(id int64) (sql.Result, error) {
	insertStmt := `
	DELETE FROM chapters WHERE book_key IN (SELECT book_id FROM books
		where book_id=$1);
	`
	result, e := db.Exec(insertStmt, id)
	if e != nil {
		println(e.Error())
		return result, e
	}

	return result, e
}
