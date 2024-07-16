package handler

import (
	"fmt"
	"net/http"

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

	fmt.Println("debug 3")
	c.JSON(http.StatusOK, gin.H{
		"message": "your book was created :)",
	})
}
