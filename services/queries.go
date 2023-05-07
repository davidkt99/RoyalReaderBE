package services

import (
	"fmt"
	"net/http"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/gin-gonic/gin"
)

func GetChapter(c *gin.Context) {
	id := URI{}
	// binding to URI
	if err := c.BindUri(&id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, db.QueryChapterById(id.Id))
}

func GetBookChapters(c *gin.Context) {
	id := URI{}
	// binding to URI
	if err := c.BindUri(&id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(id.Id)
	c.IndentedJSON(http.StatusOK, db.QueryAllChaptersOfBook(id.Id))
}

func GetBook(c *gin.Context) {
	id := URI{}
	// binding to URI
	if err := c.BindUri(&id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, db.QueryBookById(id.Id))
}

func GetAllChapterNameId(c *gin.Context) {
	id := URI{}
	// binding to URI
	if err := c.BindUri(&id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, db.QueryChapterNameAndIdsByBookId(id.Id))
}

func SayHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.QueryAllBooksWithChapterCounts())
}

func GetUpdates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.QueryWeeklyChapterUpdate())
}
