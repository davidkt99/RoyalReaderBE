package db

import (
	"fmt"

	"github.com/davidkt99/RoyalReaderBE/models"
)

func QueryWeeklyChapterUpdate() []models.ChapterUpdate {
	insertStmt := `
	select book_id, chapter_id, chapter_name, book_name, image_url, day_of_week
	from weekly_updates
	join books on weekly_updates.book_key=books.book_id
	join chapters on weekly_updates.chapter_key=chapters.chapter_id`

	var chapterUpdates []models.ChapterUpdate

	rows, e := db.Query(insertStmt)
	CheckError(e)

	defer rows.Close()
	for rows.Next() {
		update := models.ChapterUpdate{}
		e := rows.Scan(&update.BookId, &update.ChapterId, &update.ChapterName, &update.BookName, &update.ImageUrl, &update.WeekDay)
		CheckError(e)

		chapterUpdates = append(chapterUpdates, update)
	}

	//* Rows returned empty
	if len(chapterUpdates) < 1 {
		fmt.Println("No Updates Found")
		return []models.ChapterUpdate{}
	} else {
		fmt.Println("Queried # of Updates: ", len(chapterUpdates))
	}

	return chapterUpdates
}
