package repository

import (
	"github.com/oliverperboni/GoApi/schemas"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) schemas.BookRepository {
	return &BookRepositoryImpl{DB: db}
}

func (r *BookRepositoryImpl) CreateBook(book *schemas.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepositoryImpl) GetBookByID(id uint) (*schemas.Book, error) {
	var book schemas.Book
	err := r.DB.First(&book, id).Error
	return &book, err
}

func (r *BookRepositoryImpl) GetBookByName(name string) (*schemas.Book, error) {
	var book schemas.Book
	err := r.DB.First(&book, "name = ?", name).Error
	return &book, err
}

func (r *BookRepositoryImpl) UpdateBook(book *schemas.Book) error {
	return r.DB.Save(book).Error
}

func (r *BookRepositoryImpl) DeleteBook(id uint) error {
	return r.DB.Delete(&schemas.Book{}, id).Error
}
