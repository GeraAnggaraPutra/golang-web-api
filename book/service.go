package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(id int, book BookUpdate) (Book, error)
	Delete(id int) error
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
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Author:      bookRequest.Author,
		Rating:      int(rating),
		Description: bookRequest.Description,
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(id int, bookUpdate BookUpdate) (Book, error) {
	book, err := s.repository.FindByID(id)

	price, _ := bookUpdate.Price.Int64()
	rating, _ := bookUpdate.Rating.Int64()

	book.Title = bookUpdate.Title
	book.Price = int(price)
	book.Author = bookUpdate.Author
	book.Rating = int(rating)
	book.Description = bookUpdate.Description

	updateBook, err := s.repository.Update(book)
	return updateBook, err
}

func (s *service) Delete(id int) error {
	book, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	err = s.repository.Delete(book)
	return err
}
