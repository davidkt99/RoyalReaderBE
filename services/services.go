package services

import (
	"github.com/gin-gonic/gin"
)

func StartServices() {
	router := gin.Default()
	router.GET("/hello", SayHello)

	router.Run("localhost:8080")
}
