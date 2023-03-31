package services

import (
	"github.com/gin-gonic/gin"
)

func StartServices() {
	router := gin.Default()

	// Queries
	router.GET("/hello", SayHello)
	router.GET("/book/chapters/:id", GetBookChapters)
	router.GET("/book/:id", GetBook)
	router.GET("/chapter/:id", GetChapter)
	router.GET("/chapter/all/nameId/:id", GetAllChapterNameId)
	router.GET("/books", GetBooks)
	router.GET("/user/login", Login)

	// Mutations
	router.POST("/download/book/", MutationDownloadBook)
	router.POST("/user/new/", MutationNewUser)

	router.Run("0.0.0.0:8080")
}
