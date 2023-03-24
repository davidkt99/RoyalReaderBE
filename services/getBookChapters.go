package services

import (
	"fmt"
	"net/http"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/gin-gonic/gin"
)

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
