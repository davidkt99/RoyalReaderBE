package models

type Chapter struct {
	Id      int64  `json:"id"`
	BookId  int64  `json:"bookId"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Read    bool   `json:"read"`
}
