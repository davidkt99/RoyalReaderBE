package services

import (
	"net/http"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/davidkt99/RoyalReaderBE/models"
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

func MutationNewUser(c *gin.Context) {
	user := models.User{}
	// binding to URI
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, db.InsertUser(user))
}
