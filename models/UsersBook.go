package models

type UserBooks struct {
	Id          int64  `json:"id"`
	CurrMaxChap int    `json:"currMaxChap"`
	BookId      string `json:"bookId"`
}
