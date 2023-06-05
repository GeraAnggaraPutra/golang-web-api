package db

import (
	"fmt"
	"golang-web-api/book"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-web-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	fmt.Println("Database connection success")
	db.AutoMigrate(&book.Book{})
	return db
}
