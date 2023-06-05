package main

import (
	"golang-web-api/book"
	"golang-web-api/db"
	"golang-web-api/handler"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {

	bookRepository := book.NewRepository(db.Init())
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	// v1.GET("/book/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.GET("/send-mail/:subject/:body", handler.SendMail)

	// connect database
	v1.POST("/book", bookHandler.CreateBook)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/book/:id", bookHandler.GetBook)
	v1.PUT("/book/:id", bookHandler.UpdateBook)
	v1.DELETE("/book/:id", bookHandler.DeleteBook)

	v2 := router.Group("/v2")
	v2.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"content": "Ini adalah API Versioning v2",
		})
	})

	router.Run("localhost:8080")
}