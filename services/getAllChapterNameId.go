package services

import (
	"net/http"

	"github.com/davidkt99/RoyalReaderBE/db"
	"github.com/gin-gonic/gin"
)

func GetAllChapterNameId(c *gin.Context) {
	id := URI{}
	// binding to URI
	if err := c.BindUri(&id); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, db.QueryChapterNameAndIdsByBookId(id.Id))
}
