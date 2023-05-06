package scraper

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/davidkt99/RoyalReaderBE/models"
	"github.com/davidkt99/RoyalReaderBE/util"
)

func ScrapeBook(url string) (models.Book, []models.Chapter) {
	//*	Id
	id := util.BookIdFromUrl(url)

	//*	Page
	page := string(util.Download(util.ROYAL_ROAD_URL_FICTION + strconv.FormatInt(id, 10)))

	//*	Title
	title := cutTitle(page)

	//*	Author
	author := cutAuthor(page)

	//*	ImageUrl
	imageUrl := findImageUrl(page)

	//*	Chapter Urls
	chapterTable := cutChapters(page)

	matches := findAllChapterUrls(chapterTable)

	var chapters []models.Chapter
	for _, chapter := range matches {
		chapters = append(chapters, ScrapeChapter(util.ROYAL_ROAD_URL+string(chapter[1])))
	}

	book := models.Book{Id: id, Name: title, Author: author, ImageUrl: imageUrl}

	return book, chapters
}

func findImageUrl(page string) string {
	_, imgClassCut, _ := strings.Cut(string(page), "<img class=\"thumbnail inline-block\" data-type=\"avatar\"")
	_, imgCut, _ := strings.Cut(string(imgClassCut), "src=\"")
	imgUrl, _, _ := strings.Cut(string(imgCut), "\"></img>")
	return imgUrl
}

func findAllChapterUrls(chapterTable string) [][][]byte {
	re := regexp.MustCompile(`data-url="([^"]*)"`)
	match := re.FindAllSubmatch([]byte(chapterTable), -1)
	return match
}

func cutChapters(page string) string {
	_, chapterCut, _ := strings.Cut(string(page), "<table")
	chapterTable, _, _ := strings.Cut(string(chapterCut), "</table>")
	return chapterTable
}

func cutTitle(page string) string {
	_, titleCut, _ := strings.Cut(string(page), "<title>")
	titleText, _, _ := strings.Cut(string(titleCut), " | Royal Road</title>")
	return titleText
}

func cutAuthor(page string) string {
	_, profileCut, _ := strings.Cut(string(page), "<a href=\"/profile")
	_, authorCut, _ := strings.Cut(string(profileCut), "\" class=\"font-white\">")
	author, _, _ := strings.Cut(string(authorCut), "</a>")
	return author
}
