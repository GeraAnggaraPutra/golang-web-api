package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()           
	rating, _ := bookRequest.Rating.Int64()           
	book := Book {
		Title: bookRequest.Title,
		Price: int(price),
		Author: bookRequest.Author,
		Rating: int(rating),
		Description: bookRequest.Description,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}