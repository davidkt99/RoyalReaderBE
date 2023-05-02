package services

import (
	"net/http"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/gin-gonic/gin"
)

func MutationDownloadBook(c *gin.Context) {
	url := BODY_URL{}
	// binding to URI
	if err := c.BindJSON(&url); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, DownloadBook(url.Url))
}

func MutationDeleteBook(c *gin.Context) {
	id := URI{}
	// binding to URI
	if err := c.BindUri(&id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, db.DeleteBook(id.Id))
}
