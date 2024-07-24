package services

import (
	"github.com/oliverperboni/GoApi/repository"
	"github.com/oliverperboni/GoApi/schemas"
)

type ListService struct {
	repo repository.ListRepository
}

func CreateListService(r repository.ListRepository) *ListService {
	return &ListService{repo: r}
}

func (l *ListService) CreateList(list *schemas.List) error {
	return l.repo.CreateList(list)
}

func (l *ListService) AddBookToListList(bookID uint, listID uint) error {
	return l.repo.AddBookToList(bookID, listID)
}
func (l *ListService) RemoveBookToList(bookID uint, listID uint) error {
	return l.repo.RemoveBookToList(bookID, listID)
}
func (l *ListService) SeachBookToList(bookID uint, listID uint) (schemas.Book, error) {
	return l.repo.SeachBookToList(bookID, listID)

}
func (l *ListService) DeleteList(list *schemas.List) error {
	return l.repo.DeleteList(list)

}

func (l *ListService) GetAllBooksList(userID uint, listID uint) ([]schemas.Book, error) {
	return l.repo.GetAllBooksList(userID, listID)
}

func (l *ListService) GetAllUserList(userID uint) ([]schemas.List, error) {
	return l.repo.GetAllUserList(userID)
}
