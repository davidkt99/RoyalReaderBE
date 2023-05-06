package models

type BookWithChapterCount struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Author        string `json:"author"`
	ImageUrl      string `json:"imageUrl"`
	NumOfChapters int    `json:"numOfChap"`
}
