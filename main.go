package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name": "Gera Anggara Putra",
	// 		"bio":  "Software Engineer",
	// 	})
	// })

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)

	router.Run("localhost:8080")
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Gera Anggara Putra",
		"bio":  "Software Engineer",
	})
}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":   "Hello World",
		"content": "Learning Golang Web API",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": title,
	})
}

func queryHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
		"title": title,
	})
}
