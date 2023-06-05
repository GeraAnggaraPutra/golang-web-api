package book

import "encoding/json"

type BookUpdate struct {
	Title       string      `json:"title" form:"title"`
	Price       json.Number `json:"price" form:"price" binding:"number"`
	Author      string      `json:"author" form:"author"`
	Rating      json.Number `json:"rating" form:"rating" binding:"number"`
	Description string      `json:"description" form:"description"`
}