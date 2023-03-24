package services

import (
	"github.com/gin-gonic/gin"
)

func StartServices() {
	router := gin.Default()
	router.GET("/hello", SayHello)
	router.GET("/chapters/:id", GetBookChapters)
	router.GET("/book/:id", GetBook)
	router.GET("/chapter/:id", GetChapter)
	router.GET("/chapter/all/nameId/:id", GetAllChapterNameId)

	router.Run("localhost:8080")
}
