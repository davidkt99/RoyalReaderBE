package scraper

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/davidkt99/RoyalReaderBE/util"
)

func ScrapeBook(url string) {
	page := util.Download(url)

	//*	Title
	_, titleCut, _ := strings.Cut(string(page), "<title>")
	titleText, _, _ := strings.Cut(string(titleCut), " | Royal Road</title>")

	fmt.Println(titleText)

	//*	Chapter Urls
	_, chapterCut, _ := strings.Cut(string(page), "<table")
	chapterTable, _, _ := strings.Cut(string(chapterCut), "</table>")

	re := regexp.MustCompile(`href="([^"]*)"`)
	match := re.FindAllSubmatch([]byte(chapterTable), -1)

	//* Test Output
	writeErr := ioutil.WriteFile("output.txt", match[0][1], 0644)
	if writeErr != nil {
		panic(writeErr)
	}
	fmt.Println(string(match[0][1]))
}
