package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       string
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
