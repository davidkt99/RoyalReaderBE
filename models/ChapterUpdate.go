package models

type ChapterUpdate struct {
	BookId      int64  `json:"bookId"`
	ChapterId   int64  `json:"chapterId"`
	BookName    string `json:"bookName"`
	ChapterName string `json:"chapterName"`
	WeekDay     int64  `json:"weekDay"`
	ImageUrl    string `json:"imageUrl"`
}
