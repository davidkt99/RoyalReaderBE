package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World!")
}
