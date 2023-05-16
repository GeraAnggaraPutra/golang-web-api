package book


type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Price    any    `json:"price" binding:"required,number"`
	Subtitle string `json:"sub_title"` // var Subtitle dipakai untuk menangkap json yg namanya sub_title
}
