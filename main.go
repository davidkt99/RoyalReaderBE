package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	scrapChapter("https://www.royalroad.com/fiction/65359/ashes-of-heaven-book-one/chapter/1138407/chapter-24-the-great-dao-of-scamming")

}

func download(url string) []byte {
	resp, loadErr := http.Get(url)
	if loadErr != nil {
		fmt.Println("Error: ", loadErr)
		os.Exit(1)
	}
	// time.Sleep(1000 * time.Millisecond)

	// io.Copy(os.Stdout, resp.Body)

	dataInBytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("Error: ", readErr)
		os.Exit(1)
	}

	return dataInBytes
}

func scrapChapter(url string) {
	page := download(url)

	_, cut, _ := strings.Cut(string(page), "<div class=\"chapter-inner chapter-content\">")
	chapter, _, _ := strings.Cut(string(cut), "</div>")

	writeErr := ioutil.WriteFile("output.txt", []byte(chapter), 0644)
	if writeErr != nil {
		panic(writeErr)
	}
	fmt.Println(string(chapter))
}
