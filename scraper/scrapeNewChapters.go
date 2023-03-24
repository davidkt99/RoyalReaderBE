package scraper

import (
	"regexp"
	"strings"

	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func ScrapeNewChapters(url string, oldChapterNum int) []models.Chapter {
	page := util.Download(url)

	//*	Chapter Urls
	_, chapterCut, _ := strings.Cut(string(page), "<table")
	chapterTable, _, _ := strings.Cut(string(chapterCut), "</table>")

	re := regexp.MustCompile(`data-url="([^"]*)"`)
	match := re.FindAllSubmatch([]byte(chapterTable), -1)

	chapterUrls := match[oldChapterNum:]

	var chapters []models.Chapter
	for _, chapter := range chapterUrls {
		chapters = append(chapters, ScrapeChapter("https://www.royalroad.com"+string(chapter[1])))
	}

	return chapters
}
