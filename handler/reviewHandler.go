package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/services"
)

type reviewHandler struct {
	service services.ReviewService
}

func CreateReviewHandler(s services.ReviewService) reviewHandler {
	return reviewHandler{service: s}
}

func GetBookReview(c *gin.Context) {

}

func PostBookReview(c *gin.Context) {

}

func DeleteBookReview(c *gin.Context) {

}

func PutBookReview(c *gin.Context) {

}
