package scraper

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/davidkt99/RoyalReaderBE/util"
)

func ScrapeChapter(url string) {
	page := util.Download(url)

	_, cut, _ := strings.Cut(string(page), "<div class=\"chapter-inner chapter-content\">")
	chapter, _, _ := strings.Cut(string(cut), "</div>")

	writeErr := ioutil.WriteFile("output.txt", []byte(chapter), 0644)
	if writeErr != nil {
		panic(writeErr)
	}
	fmt.Println(string(chapter))
}
