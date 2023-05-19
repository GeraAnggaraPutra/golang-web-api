package book

import "encoding/json"

type BookUpdate struct {
	Title       string      `json:"title"`
	Price       json.Number `json:"price" binding:"number"`
	Author      string      `json:"author"`
	Rating      json.Number `json:"rating" binding:"number"`
	Description string      `json:"description"`
}