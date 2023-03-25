package services

import (
	"github.com/gin-gonic/gin"
)

func StartServices() {
	router := gin.Default()
	router.GET("/hello", SayHello)
	router.GET("/book/chapters/:id", GetBookChapters)
	router.GET("/book/:id", GetBook)
	router.GET("/chapter/:id", GetChapter)
	router.GET("/chapter/all/nameId/:id", GetAllChapterNameId)
	router.GET("/books", GetBooks)

	router.POST("/download/book/", MutationDownloadBook)

	router.Run("0.0.0.0:8080")
}
