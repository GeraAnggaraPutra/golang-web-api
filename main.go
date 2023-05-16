package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-api/handler"
)

func main() {
	router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name": "Gera Anggara Putra",
	// 		"bio":  "Software Engineer",
	// 	})
	// })

	router.GET("/", handler.RootHandler)
	router.GET("/hello", handler.HelloHandler)
	router.GET("/books/:id/:title", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	router.POST("/books", handler.PostBooksHandler)

	router.Run("localhost:8080")
}
