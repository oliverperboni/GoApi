package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoTomekeeper/handler"
)

// Routes:
//
//   - GET /book: Calls h.GetBook to retrieve all books.
//     Response: 200 OK with a list of books in JSON format, or 400 Bad Request if there's an error.
//
//   - GET /book/id/:id: Calls h.GetBookById to retrieve a specific book by its ID.
//     Expected Input: `id` as a URL parameter.
//     Response: 200 OK with the book in JSON format, or 400 Bad Request if the book is not found.
//
//   - GET /book/name/:name: Calls h.GetBookByName to retrieve a specific book by its name.
//     Expected Input: `name` as a URL parameter.
//     Response: 200 OK with the book in JSON format, or 400 Bad Request if the book is not found.
//
//   - POST /book: Calls h.PostBook to create a new book.
//     Expected Input: Book data in the request body as JSON.
//     Response: 200 OK with a success message, or 400 Bad Request/500 Internal Server Error if there's an error.
//
//   - PUT /book: Calls h.PutBook to update an existing book.
//     Expected Input: Updated book data in the request body as JSON.
//     Response: 200 OK with a success message, or 400 Bad Request/500 Internal Server Error if there's an error.
//
//   - DELETE /book: Calls h.DeleteBook to delete a specific book.
//     Expected Input: Book data in the request body as JSON.
//     Response: 200 OK with a success message, or 400 Bad Request/500 Internal Server Error if there's an error.
func setupBookRoutes(v *gin.RouterGroup, h *handler.BookHadler) {
	v.GET("/book", h.GetBook)

	v.GET("/book/id/:id", h.GetBookById)

	v.GET("/book/name/:name", h.GetBookByName)

	v.POST("/book", h.PostBook)

	v.PUT("/book", h.PutBook)

	v.DELETE("/book", h.DeleteBook)
}

//   - POST /lists: Calls h.PostList to create a new list.
//     Expected Input: List data in the request body as JSON.
//     Response: 200 OK if the list is created successfully, otherwise an error response.
//
//   - POST /lists/AddBook: Calls h.PostBookList to add a book to a specific list.
//     Expected Input: `bookID` and `listID` as query parameters.
//     Response: 200 OK if the book is added successfully, otherwise an error response.
//
//   - DELETE /lists: Calls h.DeleteList to delete a specific list.
//     Expected Input: List data in the request body as JSON.
//     Response: 200 OK if the list is deleted successfully, otherwise an error response.
//
//   - DELETE /list/book/: Calls h.DeleteBookList to remove a book from a specific list.
//     Expected Input: `bookID` and `listID` as query parameters.
//     Response: 200 OK if the book is removed successfully, otherwise an error response.
//
//   - GET /list/book: Calls h.GetBookList to search for a specific book in a list.
//     Expected Input: `bookID` and `listID` as query parameters.
//     Response: 200 OK with book details if found, otherwise 404 Not Found.
//
//   - GET /list/user: Calls h.GetAllUserList to retrieve all lists for a specific user.
//     Expected Input: `userID` as a query parameter.
//     Response: 200 OK with a list of user lists, otherwise 400 Bad Request.
//
//   - GET /list/allBooks: Calls h.GetAllBookList to retrieve all books from a specific list for a user.
//     Expected Input: `userID` and `listID` as query parameters.
//     Response: 200 OK with a list of books in JSON format, otherwise 400 Bad Request.
func setupListRoutes(v *gin.RouterGroup, h *handler.ListHandler) {

	v.POST("/lists", h.PostList)

	v.POST("/lists/AddBook", h.PostBookList)

	v.DELETE("/lists", h.DeleteList)

	v.DELETE("/list/book/", h.DeleteBookList)

	v.GET("/list/book", h.GetBookList)

	v.GET("/list/user", h.GetAllUserList)

	v.GET("/list/allBooks", h.GetAllBookList)
}

// setupReviewRoutes initializes the routes for review-related endpoints.
//
// Parameters:
// - v: a router group from Gin where the routes will be added. Typically, this might be a versioned group such as "v1".
// - h: an instance of ReviewHandler that contains the handler methods for the review endpoints.
func setupReviewRoutes(v *gin.RouterGroup, h *handler.ReviewHandler) {

	// GET endpoint to fetch reviews for a specific book.
	// The book ID is provided as a URL parameter.
	// Example request: GET /books/1/reviews
	v.GET("/books/:book_id/reviews", h.GetBookReview)

	// POST endpoint to create a new review.
	// The review data is provided in the request body.
	// Example request: POST /reviews
	v.POST("/reviews", h.PostBookReview)

	// DELETE endpoint to delete a review.
	// The review ID is provided as a URL parameter.
	// Example request: DELETE /reviews/1
	v.DELETE("/reviews/:review_id", h.DeleteBookReview)

	// PUT endpoint to update an existing review.
	// The review ID is provided as a URL parameter, and the updated review data is provided in the request body.
	// Example request: PUT /reviews/1
	v.PUT("/reviews/:review_id", h.PutBookReview)

}
