package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-web-api/book"
	"golang-web-api/handler"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/golang-web-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	fmt.Println("Database connection success")
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	books, err := bookRepository.FindAll()
	for _, book := range books {
		fmt.Println("=========================")
		fmt.Println("Title  : ", book.Title)
		fmt.Println("Author : ", book.Author)
	}

	// book, err := bookRepository.FindByID(2)
	// fmt.Println("Title  : ", book.Title)
	// fmt.Println("Author : ", book.Author)

	// book := book.Book{
	// 	Title: "Manusia harimau",
	// 	Author: "Person",
	// 	Price: 560000,
	// 	Rating: 4,
	// 	Description: "Sungguh aneh manusia berubah jadi harimau",
	// }
	// bookRepository.Create(book)


//////////////////////////////////////////////////////////////////////


	// CREATE
	// book := book.Book{}
	// book.Title = "Quantum theory"
	// book.Price = 230000
	// book.Author = "Kylian"
	// book.Rating = 4
	// book.Description = "Everywhere is quantum."

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("==========================")
	// }

	// READ
	// var books []book.Book
	// err = db.Debug().Where("id = ?", 1).Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }
	// for _, b := range books {
	// 	fmt.Println("Title       :", b.Title)
	// 	fmt.Println("Price       :", b.Price)
	// 	fmt.Println("Author      :", b.Author)
	// 	fmt.Println("Rating      :", b.Rating)
	// 	fmt.Println("Description :", b.Description)
	// }

	// UPDATE
	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("==========================")
	// }
	// book.Title = "Doraemon"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error updating book record")
	// 	fmt.Println("==========================")
	// }

	// DELETE
	// var book book.Book
	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("=========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=========================")
	// }
	// err = db.Delete(book).Error
	// if err != nil {
	// 	fmt.Println("=========================")
	// 	fmt.Println("Error delete book record")
	// 	fmt.Println("=========================")
	// }

	router := gin.Default()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name": "Gera Anggara Putra",
	// 		"bio":  "Software Engineer",
	// 	})
	// })

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	v2 := router.Group("/v2")

	v2.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"content": "Ini adalah API Versioning v2",
		})
	})

	router.Run("localhost:8080")
}
