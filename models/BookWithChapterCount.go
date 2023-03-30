package models

type BookWithChapterCount struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	NumOfChapters int    `json:"numOfChap"`
}
