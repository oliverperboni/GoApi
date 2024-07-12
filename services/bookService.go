package services

import "github.com/oliverperboni/GoApi/schemas"

type BookService struct {
	Repo schemas.BookRepository
}

func NewBookService(repo schemas.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (s *BookService) CreateBook(user *schemas.Book) error {
	return s.Repo.CreateBook(user)
}

func (s *BookService) GetBookByID(id uint) (*schemas.Book, error) {
	return s.Repo.GetBookByID(id)
}

func (s *BookService) GetBookByName(name string) (*schemas.Book, error) {
	return s.Repo.GetBookByName(name)
}

func (s *BookService) UpdateBook(user *schemas.Book) error {
	return s.Repo.UpdateBook(user)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.Repo.DeleteBook(id)
}
