package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoTomekeeper/schemas"
	"github.com/oliverperboni/GoTomekeeper/services"
	"github.com/oliverperboni/GoTomekeeper/utils"
)

// BookHadler manages HTTP requests related to books, using the provided BookService.
type BookHadler struct {
	service services.BookService
}

// CreateBookHandler initializes a new BookHadler with the given BookService.
func CreateBookHandler(s *services.BookService) BookHadler {
	return BookHadler{service: *s}
}

// PostBook handles the creation of a new book.
// It expects the book data in the request body as JSON.
// On success, it returns a 200 OK response with a success message.
// On failure, it returns a 400 Bad Request or 500 Internal Server Error with an error message.
func (h *BookHadler) PostBook(c *gin.Context) {
	var bookJSON schemas.Book
	if err := c.ShouldBindJSON(&bookJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createErr := h.service.PostBook(bookJSON); createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your book was created :)",
	})
}

// GetBook retrieves all books from the database.
// It returns a list of books in JSON format with a 200 OK response.
// On failure, it returns a 400 Bad Request with an error message.
func (h *BookHadler) GetBook(c *gin.Context) {
	var books []schemas.Book
	var booksJSON []schemas.BookJSON
	books, err := h.service.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, book := range books {
		booksJSON = append(booksJSON, utils.BookConvert(book))
	}

	c.JSON(http.StatusOK, booksJSON)
}

// GetBookById retrieves a specific book by its ID.
// It expects the ID as a URL parameter.
// On success, it returns the book in JSON format with a 200 OK response.
// On failure, it returns a 400 Bad Request with an error message.
func (h *BookHadler) GetBookById(c *gin.Context) {
	str := c.Param("id")
	id, _ := strconv.Atoi(str)
	bookID := uint(id)
	book, err := h.service.GetBookByID(bookID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookJSON := utils.BookConvert(*book)

	c.JSON(http.StatusOK, bookJSON)
}

// GetBookByName retrieves a specific book by its name.
// It expects the name as a URL parameter.
// On success, it returns the book in JSON format with a 200 OK response.
// On failure, it returns a 400 Bad Request with an error message.
func (h *BookHadler) GetBookByName(c *gin.Context) {
	name := c.Param("name")

	book, err := h.service.GetBookByName(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookJSON := utils.BookConvert(*book)

	c.JSON(http.StatusOK, bookJSON)
}

// PutBook handles the update of an existing book.
// It expects the updated book data in the request body as JSON.
// On success, it returns a 200 OK response with a success message.
// On failure, it returns a 400 Bad Request or 500 Internal Server Error with an error message.
func (h *BookHadler) PutBook(c *gin.Context) {
	var bookJSON schemas.Book
	if err := c.ShouldBindJSON(&bookJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createErr := h.service.UpdateBook(&bookJSON); createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your book was updated :)",
	})
}

// DeleteBook handles the deletion of a book.
// It expects the book data in the request body as JSON.
// On success, it returns a 200 OK response with a success message.
// On failure, it returns a 400 Bad Request or 500 Internal Server Error with an error message.
func (h *BookHadler) DeleteBook(c *gin.Context) {
	var bookJSON schemas.Book
	if err := c.ShouldBindJSON(&bookJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if createErr := h.service.DeleteBook(&bookJSON); createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": createErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your book was deleted :)",
	})
}
