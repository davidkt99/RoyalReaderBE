package services

import (
	"github.com/gin-gonic/gin"
)

func StartServices() {
	router := gin.Default()
	router.GET("/hello", SayHello)
	router.GET("/chapters/:id", GetBookChapters)
	router.GET("/book/:id", GetBook)

	router.Run("localhost:8080")
}
