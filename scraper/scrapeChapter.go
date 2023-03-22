package scraper

import (
	"strings"

	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func ScrapeChapter(url string) models.Chapter {
	page := util.Download(url)

	_, cutTitle, _ := strings.Cut(string(page), `<h1 style="margin-top: 10px" class="font-white">`)
	title, _, _ := strings.Cut(string(cutTitle), "</h1>")

	_, cut, _ := strings.Cut(string(page), "<div class=\"chapter-inner chapter-content\">")
	content, _, _ := strings.Cut(string(cut), "</div>")

	chapter := models.Chapter{Name: title, Content: content}

	return chapter
}
