package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/schemas"
	"github.com/oliverperboni/GoApi/services"
	"github.com/oliverperboni/GoApi/utils"

	"strconv"
)

type ListHandler struct {
	service services.ListService
}

func CreateNewListHandler(s *services.ListService) ListHandler {
	return ListHandler{service: *s}
}

func (l *ListHandler) GetAllBookList(c *gin.Context) {
	// get all books of a list
	//input: userID and listID
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
	c.JSON(http.StatusOK, books)
}

func (l *ListHandler) GetAllUserList(c *gin.Context) {
	// get all lists of an user
	//input: userID
}

func (l *ListHandler) GetBookList(c *gin.Context) {
	// search a book in a list
	// input: bookID and listID
}

func (l *ListHandler) DeleteList(c *gin.Context) {
	//delete a list
	// input: list
	var list schemas.List
	if err := c.ShouldBindJSON(&list); err != nil {
		// Handle JSON binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func (l *ListHandler) PostList(c *gin.Context) {
	//create new list
	// input: List
	var list schemas.List
	if err := c.ShouldBindJSON(&list); err != nil {
		// Handle JSON binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createErr := l.service.CreateList(&list); createErr != nil {
		// Handle error from service.PostBook
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your list was created :)",
	})


}

func (l *ListHandler) PostBookList(c *gin.Context) {
	// add book to the list
	//input bookID and listID
	book, _ := strconv.Atoi(c.Query("bookID"))
	list, _ := strconv.Atoi(c.Query("listID"))
	bookID := uint(book)
	listID := uint(list)
	err := l.service.AddBookToListList(bookID,listID)
	if  err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err })
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "book added :) "})
	return


}

func (l *ListHandler) DeleteBookList(c *gin.Context) {
	// remove book to the list
	//input bookID and listID
}
