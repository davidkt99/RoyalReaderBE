package scraper

import (
	"regexp"
	"strings"

	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func ScrapeBook(url string) (string, string, []models.Chapter) {
	page := util.Download(url)

	//*	Title
	_, titleCut, _ := strings.Cut(string(page), "<title>")
	titleText, _, _ := strings.Cut(string(titleCut), " | Royal Road</title>")

	//*	Chapter Urls
	_, chapterCut, _ := strings.Cut(string(page), "<table")
	chapterTable, _, _ := strings.Cut(string(chapterCut), "</table>")

	re := regexp.MustCompile(`data-url="([^"]*)"`)
	match := re.FindAllSubmatch([]byte(chapterTable), -1)

	var chapters []models.Chapter
	for _, chapter := range match {
		chapters = append(chapters, ScrapeChapter("https://www.royalroad.com"+string(chapter[1])))
	}

	return titleText, url, chapters
}
