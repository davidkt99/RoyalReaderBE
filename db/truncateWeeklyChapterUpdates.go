package db

import "database/sql"

func TruncateWeeklyChapterUpdates() sql.Result {
	insertStmt := `
	TRUNCATE TABLE weekly_updates;
	`
	result, e := db.Exec(insertStmt)
	if e != nil {
		println(e.Error())
	}

	numOfRowsAffected, _ := result.RowsAffected()

	println("Truncated Weekly Chapter Updates: ", numOfRowsAffected)

	return result
}
