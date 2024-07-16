package repository

import (
	"github.com/oliverperboni/GoApi/schemas"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func CreateBookRepository(db *gorm.DB) BookRepository {
	return BookRepository{DB: db}
}

func (r *BookRepository) CreateBook(book *schemas.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) GetBookByID(id uint) (*schemas.Book, error) {
	var book schemas.Book
	err := r.DB.First(&book, "Id = ?", id).Error
	return &book, err
}

func (r *BookRepository) GetBookByName(name string) (*schemas.Book, error) {
	var book schemas.Book
	err := r.DB.Where("Name = ?", name).First(&book).Error
	return &book, err

}

func (r *BookRepository) UpdateBook(book *schemas.Book) error {
	return r.DB.Save(&book).Error
}

func (r *BookRepository) DeleteBook(book *schemas.Book) error {
	return r.DB.Delete(&book).Where("id=?", book.Id).Error
}

func (r *BookRepository) GetBooks() ([]schemas.Book, error) {
	var books []schemas.Book
	err := r.DB.Find(&books).Error
	return books, err
}
