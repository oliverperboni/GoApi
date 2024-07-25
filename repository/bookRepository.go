package repository

import (
	"github.com/oliverperboni/GoApi/schemas"
	"gorm.io/gorm"
)

// BookRepository provides methods to interact with the books in the database.
type BookRepository struct {
	DB *gorm.DB
}

// CreateBookRepository initializes a new BookRepository with the given GORM database connection.
func CreateBookRepository(db *gorm.DB) BookRepository {
	return BookRepository{DB: db}
}

// CreateBook adds a new book to the database.
// It accepts a pointer to a schemas.Book and attempts to create a new record.
// Returns an error if the operation fails.
func (r *BookRepository) CreateBook(book *schemas.Book) error {
	return r.DB.Create(book).Error
}

// GetBookByID retrieves a book from the database by its ID.
// It accepts the book ID and returns a pointer to the book and an error if the operation fails.
func (r *BookRepository) GetBookByID(id uint) (*schemas.Book, error) {
	var book schemas.Book
	err := r.DB.First(&book, "Id = ?", id).Error
	return &book, err
}

// GetBookByName retrieves a book from the database by its name.
// It accepts the book name and returns a pointer to the book and an error if the operation fails.
func (r *BookRepository) GetBookByName(name string) (*schemas.Book, error) {
	var book schemas.Book
	err := r.DB.Where("Name = ?", name).First(&book).Error
	return &book, err
}

// UpdateBook updates an existing book in the database with new data.
// It accepts a pointer to a schemas.Book and attempts to save the updated record.
// Returns an error if the operation fails.
func (r *BookRepository) UpdateBook(book *schemas.Book) error {
	return r.DB.Save(book).Error
}

// DeleteBook removes a book from the database.
// It accepts a pointer to a schemas.Book and attempts to delete the record by its ID.
// Returns an error if the operation fails.
func (r *BookRepository) DeleteBook(book *schemas.Book) error {
	return r.DB.Delete(book, "Id = ?", book.ID).Error
}

// GetBooks retrieves all books from the database.
// It returns a slice of books and an error if the operation fails.
func (r *BookRepository) GetBooks() ([]schemas.Book, error) {
	var books []schemas.Book
	err := r.DB.Find(&books).Error
	return books, err
}
