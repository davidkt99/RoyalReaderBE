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

	//*	Chapter Urls
	_, cut, _ := strings.Cut(string(page), "<table")
	table, _, _ := strings.Cut(string(cut), "</table>")

	re := regexp.MustCompile(`href="([^"]*)"`)
	match := re.FindAllSubmatch([]byte(table), -1)

	//* Test Output
	writeErr := ioutil.WriteFile("output.txt", match[0][1], 0644)
	if writeErr != nil {
		panic(writeErr)
	}
	fmt.Println(string(match[0][1]))
}
