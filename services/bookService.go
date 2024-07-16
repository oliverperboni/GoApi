package services

import (
	"github.com/oliverperboni/GoApi/repository"
	"github.com/oliverperboni/GoApi/schemas"
)

type BookService struct {
	Repo repository.BookRepository
}

func CreateBookService(r repository.BookRepository) *BookService {
	return &BookService{Repo: r}
}

func (s *BookService) PostBook(b schemas.Book) error {
	return s.Repo.CreateBook(&b)
}
