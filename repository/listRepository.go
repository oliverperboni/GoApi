package repository

import (
	"time"

	"github.com/oliverperboni/GoApi/schemas"
	"gorm.io/gorm"
)

// ListRepository provides methods to interact with the lists and books in the database.
type ListRepository struct {
	DB *gorm.DB
}

// CreateListRepository initializes a new ListRepository with the given GORM database connection.
func CreateListRepository(db *gorm.DB) ListRepository {
	return ListRepository{DB: db}
}

// CreateList adds a new list to the database.
// It accepts a pointer to a schemas.List and attempts to create a new record.
// Returns an error if the operation fails.
func (l *ListRepository) CreateList(list *schemas.List) error {
	return l.DB.Create(list).Error
}

// AddBookToList associates a book with a list by creating a new ListBook record.
// It accepts the IDs of the book and the list, and sets the creation and update timestamps.
// Returns an error if the operation fails.
func (l *ListRepository) AddBookToList(bookID uint, listID uint) error {
	listbook := &schemas.ListBook{
		BookID:    bookID,
		ListID:    listID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return l.DB.Create(listbook).Error
}

// RemoveBookToList removes the association between a book and a list by deleting the corresponding ListBook record.
// It accepts the IDs of the book and the list.
// Returns an error if the operation fails.
func (l *ListRepository) RemoveBookToList(bookID uint, listID uint) error {
	var listbook schemas.ListBook
	l.DB.Where("Book_ID = ? AND List_ID = ?", bookID, listID).Find(&listbook)
	return l.DB.Delete(&listbook).Error
}

// UpdateList updates an existing list in the database with new data.
// It accepts a pointer to a schemas.List and attempts to save the updated record.
// Returns an error if the operation fails.
func (l *ListRepository) UpdateList(list *schemas.List) error {
	return l.DB.Save(list).Error
}

// SeachBookToList searches for a book within a list.
// It accepts the IDs of the book and the list, and returns the book if it exists in the list.
// Returns the book and an error if the operation fails or the book is not found.
func (l *ListRepository) SeachBookToList(bookID uint, listID uint) (schemas.Book, error) {
	var listbook schemas.ListBook
	var book schemas.Book
	err := l.DB.Where("Book_ID = ? AND List_ID = ?", bookID, listID).Find(&listbook).Error
	if err != nil {
		return book, err
	}

	err = l.DB.Find(&book, "Id = ?", listbook.BookID).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

// GetAllBooksList retrieves all books associated with a specific list for a user.
// It accepts the IDs of the user and the list, and returns a slice of books.
// Returns a slice of books and an error if the operation fails.
func (l *ListRepository) GetAllBooksList(userID uint, listID uint) ([]schemas.Book, error) {
	var listBook []schemas.ListBook
	var book schemas.Book
	var books []schemas.Book
	var list schemas.List

	err := l.DB.Where("User_ID = ? AND Id = ?", userID, listID).Find(&list).Error
	if err != nil {
		return books, err
	}

	err = l.DB.Find(&listBook, "list_ID = ?", list.ID).Error
	if err != nil {
		return books, err
	}
	for _, v := range listBook {
		err := l.DB.Where("Id = ?", v.BookID).Find(&book).Error
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}

// GetAllUserList retrieves all lists associated with a specific user.
// It accepts the userID and returns a slice of lists.
// Returns a slice of lists and an error if the operation fails.
func (l *ListRepository) GetAllUserList(userID uint) ([]schemas.List, error) {
	var lists []schemas.List
	err := l.DB.Find(&lists, "User_ID = ?", userID).Error
	return lists, err
}

// DeleteList removes a list from the database.
// It accepts a pointer to a schemas.List and attempts to delete the record.
// Returns an error if the operation fails.
func (l *ListRepository) DeleteList(list *schemas.List) error {
	return l.DB.Delete(list).Error
}
