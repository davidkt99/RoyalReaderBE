package scraper

import (
	"strconv"
	"strings"

	"github.com/davidkt99/RoyalReaderBE/util"
)

func ScrapeChatperNum(url string) int {
	page := util.Download(url)

	//*	Num of Chapters
	_, numCut, _ := strings.Cut(string(page), "data-chapters=\"")
	num, _, _ := strings.Cut(string(numCut), "\">")

	count, _ := strconv.Atoi(num)

	return count
}
