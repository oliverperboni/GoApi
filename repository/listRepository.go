package repository

// TODO Tests
import (
	"time"

	"github.com/oliverperboni/GoApi/schemas"
	"gorm.io/gorm"
)

type ListRepository struct {
	DB *gorm.DB
}

func CreateListRepository(db *gorm.DB) ListRepository {
	return ListRepository{DB: db}
}

func (l *ListRepository) CreateList(list *schemas.List) error {
	return l.DB.Create(list).Error
}

func (l *ListRepository) AddBookToList(bookID uint, listID uint) error {
	var listbook schemas.ListBook
	listbook.BookID = bookID
	listbook.ListID = listID
	listbook.CreatedAt = time.Now()
	return l.DB.Create(listbook).Error
}

func (l *ListRepository) RemoveBookToList(bookID uint, listID uint) error {
	var listbook schemas.ListBook
	l.DB.Where("BookID = ? AND ListID = ?", bookID, listID).Find(&listbook)
	return l.DB.Delete(listbook).Error
}

func (l *ListRepository) SeachBookToList(bookID uint, listID uint) (schemas.Book, error) {
	var listbook schemas.ListBook
	var book schemas.Book
	err := l.DB.Where("BookID = ? AND ListID = ?", bookID, listID).Find(&listbook).Error
	if err != nil {
		return book, err
	}

	err = l.DB.Find(&book, "Id = ?", listbook.BookID).Error

	if err != nil {
		return book, err
	}

	return book, err
}

func (l *ListRepository) GetAllBooksList(userID uint, listID uint) ([]schemas.ListBook, error) {

	var listBook []schemas.ListBook
	var list schemas.List

	err := l.DB.Where("UserID = ? AND Id = ?", userID, listID).Find(&list).Error

	if err != nil {
		return listBook, err
	}

	err = l.DB.Find(&listBook, "listID = ?", list.ID).Error
	if err != nil {
		return listBook, err
	}
	return listBook, err
}

func (l *ListRepository) GetAllUserList(userID uint) (schemas.List, error) {
	var list schemas.List
	err := l.DB.Find(&list, "UserID = ?", userID).Error
	return list, err
}

func (l *ListRepository) DeleteList(list *schemas.List) error {
	return l.DB.Delete(list).Error
}
