package scraper

import (
	"fmt"
	"strings"

	"github.com/davidkt99/RoyalReaderBE/util"
)

func CheckNewChapters(url string) bool {
	page := util.Download(url)

	//*	Num of Chapters
	_, numCut, _ := strings.Cut(string(page), "data-chapters=\"")
	num, _, _ := strings.Cut(string(numCut), "\">")

	fmt.Println(num)

	return false
}
