package services

import (
	"github.com/oliverperboni/GoTomekeeper/repository"
	"github.com/oliverperboni/GoTomekeeper/schemas"
)

type BookService struct {
	repo repository.BookRepository
}

func CreateBookService(r repository.BookRepository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) PostBook(b schemas.Book) error {
	return s.repo.CreateBook(&b)
}

func (s *BookService) GetBooks() ([]schemas.Book, error) {
	return s.repo.GetBooks()
}

func (s *BookService) GetBookByID(id uint) (*schemas.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *BookService) GetBookByName(name string) (*schemas.Book, error) {
	return s.repo.GetBookByName(name)
}

func (s *BookService) UpdateBook(b *schemas.Book) error {
	return s.repo.UpdateBook(b)
}

func (s *BookService) DeleteBook(b *schemas.Book) error {
	return s.repo.DeleteBook(b)
}
