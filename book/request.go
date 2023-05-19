package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Author      string      `json:"author" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
}
