package services

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServices() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
	}))

	router.GET("/hello", SayHello)
	router.GET("/book/chapters/:id", GetBookChapters)
	router.GET("/book/:id", GetBook)
	router.GET("/chapter/:id", GetChapter)
	router.GET("/chapter/all/nameId/:id", GetAllChapterNameId)
	router.GET("/books", GetBooks)
	router.GET("/updates", GetUpdates)

	router.POST("/download/book/", MutationDownloadBook)
	router.POST("/delete/:id", MutationDeleteBook)

	router.Run("0.0.0.0:8080")
}
