package services

import (
	"net/http"

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
