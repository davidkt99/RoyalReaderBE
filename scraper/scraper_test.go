package scraper

import (
	"testing"

	"github.com/davidkt99/RoyalReaderBE/models"
)

const ExpectedTitle = "A Novel Concept"
const ExpectedId = 66455
const ExpectedAuthor = "Priam"
const ExpectedImageUrl = "https://www.royalroadcdn.com/public/covers-large/a-novel-concept-aacaiedfixm.jpg?time=1680391196"
const MinExpectedChapterCount = 10

var ExpectedBook = models.Book{
	Id:       ExpectedId,
	Name:     ExpectedTitle,
	Author:   ExpectedAuthor,
	ImageUrl: ExpectedImageUrl,
}

func TestScrapeBook(t *testing.T) {
	book, chapters := ScrapeBook("https://www.royalroad.com/fiction/66455/a-novel-concept")
	if !(book.Name == ExpectedTitle) {
		t.Errorf("Expected title to be %v, got %v", ExpectedTitle, book.Name)
	}
	if !(book.Id == ExpectedId) {
		t.Errorf("Expected url to be %v, got %v", ExpectedId, book.Id)
	}
	if !(book.Author == ExpectedAuthor) {
		t.Errorf("Expected author to be %v, got %v", ExpectedAuthor, book.Author)
	}
	if !(len(chapters) >= MinExpectedChapterCount) {
		t.Errorf("Expected at least %v chapters, got %v", MinExpectedChapterCount, len(chapters))
	}
}
