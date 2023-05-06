package util

import (
	"log"
	"regexp"
	"strconv"
)

func BookIdFromUrl(url string) int64 {
	idRe := regexp.MustCompile("https://www.royalroad.com/fiction/([0-9]*)")
	idMatch := idRe.FindAllSubmatch([]byte(url), -1)
	id, err := strconv.ParseInt(string(idMatch[0][1]), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return id
}
