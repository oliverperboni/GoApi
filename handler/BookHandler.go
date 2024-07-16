package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/schemas"
	"github.com/oliverperboni/GoApi/services"
)

type BookHadler struct {
	service services.BookService
}

func CreateBookHandler(s *services.BookService) BookHadler {
	return BookHadler{service: *s}

}

func (h *BookHadler) PostBook(c *gin.Context) {
	var bookJSON schemas.Book
	if err := c.ShouldBindJSON(&bookJSON); err != nil {
		// Handle JSON binding error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createErr := h.service.PostBook(bookJSON); createErr != nil {
		// Handle error from service.PostBook
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your book was created :)",
	})
}

func (h *BookHadler) GetBook(c *gin.Context) {
	var book []schemas.Book
	book, err := h.service.Repo.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)

}
func (h *BookHadler) GetBookById(c *gin.Context) {
	//get the id param

	str := c.Param("id")
	id, _ := strconv.Atoi(str)
	bookID := uint(id)
	book, err := h.service.Repo.GetBookByID(bookID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)

}
func (h *BookHadler) GetBookByName(c *gin.Context) {
	//get the name param
}
func (h *BookHadler) PutBook(c *gin.Context) {
	//get the book from the body

}
func (h *BookHadler) DeleteBook(c *gin.Context) {
	//get the book from the body

}
