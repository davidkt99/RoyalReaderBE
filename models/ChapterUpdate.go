package models

type ChapterUpdate struct {
	BookId      int64  `json:"bookId"`
	ChapterId   string `json:"chapterId"`
	BookName    string `json:"bookName"`
	ChapterName string `json:"chapterName"`
	WeekDay     int64  `json:"weekDay"`
	Name        string `json:"name"`
	ImageUrl    string `json:"imageUrl"`
}
