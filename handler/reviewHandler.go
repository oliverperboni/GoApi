package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoTomekeeper/schemas"
	"github.com/oliverperboni/GoTomekeeper/services"
)

// ReviewHandler handles HTTP requests for review-related operations.
type ReviewHandler struct {
	service services.ReviewService
}

// CreateReviewHandler creates a new instance of ReviewHandler with the given service.
func CreateReviewHandler(s services.ReviewService) ReviewHandler {
	return ReviewHandler{service: s}
}

// GetBookReview handles the GET request to fetch reviews for a specific book.
// The book ID is retrieved from the URL parameter.
func (h ReviewHandler) GetBookReview(c *gin.Context) {
	// Parse the book ID from the URL parameter and convert it to an integer.
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		// Respond with a bad request status if the book ID is invalid.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Fetch reviews for the specified book ID using the service.
	reviews, err := h.service.ReadReviewByBook(uint(bookID))
	if err != nil {
		// Respond with an internal server error status if there is an error fetching the reviews.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the fetched reviews.
	c.JSON(http.StatusOK, reviews)
}

// PostBookReview handles the POST request to create a new review.
// The review data is retrieved from the request body.
func (h ReviewHandler) PostBookReview(c *gin.Context) {
	var review schemas.Review
	// Bind the JSON request body to the review struct.
	if err := c.ShouldBindJSON(&review); err != nil {
		// Respond with a bad request status if there is an error binding the JSON.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the review using the service.
	if err := h.service.CreateReview(&review); err != nil {
		// Respond with a bad request status if there is an error creating the review.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Respond with a success message if the review is created successfully.
	c.JSON(http.StatusOK, gin.H{"msg": "review created :)"})
}

// DeleteBookReview handles the DELETE request to delete a review.
// The review ID is retrieved from the URL parameter.
func (h ReviewHandler) DeleteBookReview(c *gin.Context) {
	// Parse the review ID from the URL parameter and convert it to an integer.
	reviewID, err := strconv.Atoi(c.Param("review_id"))
	if err != nil {
		// Respond with a bad request status if the review ID is invalid.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	// Delete the review using the service.
	if err := h.service.DeleteReview(uint(reviewID)); err != nil {
		// Respond with an internal server error status if there is an error deleting the review.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message if the review is deleted successfully.
	c.JSON(http.StatusOK, gin.H{"msg": "review deleted :)"})
}

// PutBookReview handles the PUT request to update an existing review.
// The review data is retrieved from the request body.
func (h ReviewHandler) PutBookReview(c *gin.Context) {
	var review schemas.Review
	// Bind the JSON request body to the review struct.
	if err := c.ShouldBindJSON(&review); err != nil {
		// Respond with a bad request status if there is an error binding the JSON.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the review using the service.
	if err := h.service.UpdateReview(&review); err != nil {
		// Respond with an internal server error status if there is an error updating the review.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message if the review is updated successfully.
	c.JSON(http.StatusOK, gin.H{"msg": "review updated :)"})
}
