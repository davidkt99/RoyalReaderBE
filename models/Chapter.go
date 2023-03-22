package models

type Chapter struct {
	Id      int64
	BookId  int64
	Name    string
	Content string
	Read    bool
}
