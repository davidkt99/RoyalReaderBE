package services

type URI struct {
	Id int64 `json:"id" uri:"id"`
}

type BODY_URL struct {
	Url string `json:"url" uri:"url"`
}
