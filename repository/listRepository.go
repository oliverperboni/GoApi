package repository

// TODO Implementations
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

func (l *ListRepository) RemoveBookToList(bookId uint) error {
	var listbook schemas.ListBook
	l.DB.Find(&listbook, "BookID = ?", bookId)
	return l.DB.Delete(listbook).Error
}

func (l *ListRepository) SeachBookToList(bookId uint) schemas.Book {
	return schemas.Book{}
}

func (l *ListRepository) DeleteList(list *schemas.List) error {
	return l.DB.Delete(list).Error
}
