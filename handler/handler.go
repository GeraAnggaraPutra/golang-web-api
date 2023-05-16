package handler

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang-web-api/book"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Gera Anggara Putra",
		"bio":  "Software Engineer",
	})
}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":   "Hello World",
		"content": "Learning Golang Web API",
	})
}

func BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	title := ctx.Query("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func PostBooksHandler(ctx *gin.Context) {
	var bookInput book.BookInput

	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"price":    bookInput.Price,
		"subtitle": bookInput.Subtitle,
	})
}
