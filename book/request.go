package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" form:"title" binding:"required"`
	Price       json.Number `json:"price" form:"price" binding:"required,number"`
	Author      string      `json:"author" form:"author" binding:"required"`
	Rating      json.Number `json:"rating" form:"rating" binding:"required,number"`
	Description string      `json:"description" form:"description" binding:"required"`
}
