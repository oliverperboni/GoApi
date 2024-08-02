package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/oliverperboni/GoTomekeeper/schemas"
	"github.com/oliverperboni/GoTomekeeper/services"
	"github.com/oliverperboni/GoTomekeeper/utils"
)

// ListHandler manages requests related to book lists.
type ListHandler struct {
	service services.ListService
}

// CreateNewListHandler initializes a new ListHandler with the provided ListService.
func CreateNewListHandler(s *services.ListService) ListHandler {
	return ListHandler{service: *s}
}

// GetAllBookList retrieves all books from a specific list for a user.
// It expects `userID` and `listID` as query parameters.
// Responses:
//   - 200 OK: Returns a list of books in JSON format.
//   - 400 Bad Request: Returns an error if there's an issue with fetching the books.
func (l *ListHandler) GetAllBookList(c *gin.Context) {
	user, _ := strconv.Atoi(c.Query("userID"))
	list, _ := strconv.Atoi(c.Query("listID"))
	userID := uint(user)
	listID := uint(list)
	var booksJSON []schemas.BookJSON
	books, err := l.service.GetAllBooksList(userID, listID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range books {
		booksJSON = append(booksJSON, utils.BookConvert(v))
	}
	c.JSON(http.StatusOK, booksJSON)
}

// GetAllUserList retrieves all lists for a specific user.
// It expects `userID` as a query parameter.
// Responses:
//   - 200 OK: Returns a list of lists in JSON format.
//   - 400 Bad Request: Returns an error if there's an issue with fetching the lists.
func (l *ListHandler) GetAllUserList(c *gin.Context) {
	user, _ := strconv.Atoi(c.Query("userID"))
	userID := uint(user)
	lists, err := l.service.GetAllUserList(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, lists)
}

// GetBookList searches for a specific book within a list.
// It expects `bookID` and `listID` as query parameters.
// Responses:
//   - 200 OK: Returns the book in JSON format if found.
//   - 404 Not Found: Returns an error if the book is not found in the list.
func (l *ListHandler) GetBookList(c *gin.Context) {
	book, _ := strconv.Atoi(c.Query("bookID"))
	list, _ := strconv.Atoi(c.Query("listID"))
	bookID := uint(book)
	listID := uint(list)

	findBook, err := l.service.SeachBookToList(bookID, listID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "this book does not exist in this list"})
		return
	}

	c.JSON(http.StatusOK, utils.BookConvert(findBook))
}

// DeleteList removes a list from the system.
// It expects the list data to be provided in the request body as JSON.
// Responses:
//   - 200 OK: Confirms the list was deleted.
//   - 400 Bad Request: Returns an error if there's an issue with binding the JSON data.
//   - 404 Not Found: Returns an error if the list could not be found or deleted.
func (l *ListHandler) DeleteList(c *gin.Context) {
	var list schemas.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := l.service.DeleteList(&list); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"msg": "list deleted :)"})
}

// PostList creates a new list.
// It expects the list data to be provided in the request body as JSON.
// Responses:
//   - 200 OK: Confirms the list was created.
//   - 400 Bad Request: Returns an error if there's an issue with binding the JSON data.
//   - 500 Internal Server Error: Returns an error if there's an issue creating the list.
func (l *ListHandler) PostList(c *gin.Context) {
	var list schemas.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createErr := l.service.CreateList(&list); createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your list was created :)",
	})
}

// PostBookList adds a book to a specific list.
// It expects `bookID` and `listID` as query parameters.
// Responses:
//   - 200 OK: Confirms the book was added.
//   - 500 Internal Server Error: Returns an error if there's an issue adding the book.
func (l *ListHandler) PostBookList(c *gin.Context) {
	book, _ := strconv.Atoi(c.Query("bookID"))
	list, _ := strconv.Atoi(c.Query("listID"))
	bookID := uint(book)
	listID := uint(list)
	err := l.service.AddBookToListList(bookID, listID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "book added :) "})
}

// DeleteBookList removes a book from a specific list.
// It expects `bookID` and `listID` as query parameters.
// Responses:
//   - 200 OK: Confirms the book was removed.
//   - 500 Internal Server Error: Returns an error if there's an issue removing the book.
func (l *ListHandler) DeleteBookList(c *gin.Context) {
	book, _ := strconv.Atoi(c.Query("bookID"))
	list, _ := strconv.Atoi(c.Query("listID"))
	bookID := uint(book)
	listID := uint(list)
	if err := l.service.RemoveBookToList(bookID, listID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "book was deleted from the list :)"})
}
