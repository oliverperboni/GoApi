package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoTomekeeper/handler"
	"github.com/oliverperboni/GoTomekeeper/middleware"
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
	v.GET("/book", middleware.RequireAuth, h.GetBook)

	v.GET("/book/id/:id", middleware.RequireAuth, h.GetBookById)

	v.GET("/book/name/:name", middleware.RequireAuth, h.GetBookByName)

	v.POST("/book", middleware.RequireAuth, h.PostBook)

	v.PUT("/book", middleware.RequireAuth, h.PutBook)

	v.DELETE("/book", middleware.RequireAuth, h.DeleteBook)
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

	v.POST("/lists", middleware.RequireAuth, h.PostList)

	v.POST("/lists/AddBook", middleware.RequireAuth, h.PostBookList)

	v.DELETE("/lists", middleware.RequireAuth, h.DeleteList)

	v.DELETE("/list/book/", middleware.RequireAuth, h.DeleteBookList)

	v.GET("/list/book", middleware.RequireAuth, h.GetBookList)

	v.GET("/list/user", middleware.RequireAuth, h.GetAllUserList)

	v.GET("/list/allBooks", middleware.RequireAuth, h.GetAllBookList)
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
	v.GET("/books/:book_id/reviews", middleware.RequireAuth, h.GetBookReview)

	// POST endpoint to create a new review.
	// The review data is provided in the request body.
	// Example request: POST /reviews
	v.POST("/reviews", middleware.RequireAuth, h.PostBookReview)

	// DELETE endpoint to delete a review.
	// The review ID is provided as a URL parameter.
	// Example request: DELETE /reviews/1
	v.DELETE("/reviews/:review_id", middleware.RequireAuth, h.DeleteBookReview)

	// PUT endpoint to update an existing review.
	// The review ID is provided as a URL parameter, and the updated review data is provided in the request body.
	// Example request: PUT /reviews/1
	v.PUT("/reviews/:review_id", middleware.RequireAuth, h.PutBookReview)

}

func setupUserRoutes(v *gin.RouterGroup, h *handler.UserHandler) {
	v.POST("/user", h.CreateUser)

	v.GET("/user", h.GetUser)
}
