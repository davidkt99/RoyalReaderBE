package scraper

import (
	"testing"

	"github.com/davidkt99/RoyalReaderBE/models"
)

const ExpectedTitle = "Mother of Learning"
const ExpectedId = 21220
const ExpectedAuthor = "nobody103"
const ExpectedImageUrl = "https://www.royalroadcdn.com/public/covers-full/21220-mother-of-learning.jpg?time=1637247458"
const MinExpectedChapterCount = 10

var ExpectedBook = models.Book{
	Id:       ExpectedId,
	Name:     ExpectedTitle,
	Author:   ExpectedAuthor,
	ImageUrl: ExpectedImageUrl,
}

func TestScrapeBook(t *testing.T) {
	book, chapters := ScrapeBook("https://www.royalroad.com/fiction/21220/mother-of-learning")
	if !(book.Name == ExpectedTitle) {
		t.Errorf("Expected title to be %v, got %v", ExpectedTitle, book.Name)
	}
	if !(book.Id == ExpectedId) {
		t.Errorf("Expected url to be %v, got %v", ExpectedId, book.Id)
	}
	if !(book.Author == ExpectedAuthor) {
		t.Errorf("Expected author to be %v, got %v", ExpectedAuthor, book.Author)
	}
	if !(book.ImageUrl == ExpectedImageUrl) {
		t.Errorf("Expected image url to be %v, got %v", ExpectedImageUrl, book.ImageUrl)
	}
	if !(len(chapters) >= MinExpectedChapterCount) {
		t.Errorf("Expected at least %v chapters, got %v", MinExpectedChapterCount, len(chapters))
	}
}
