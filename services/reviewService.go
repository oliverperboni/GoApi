package services

import (
	"github.com/oliverperboni/GoApi/repository"
	"github.com/oliverperboni/GoApi/schemas"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func CreateReviewService(r repository.ReviewRepository) ReviewService {
	return ReviewService{repo: r}
}

func (r *ReviewService) CreateReview(review *schemas.Review) error {
	return r.CreateReview(review)
}

func (r *ReviewService) UpdateReview(review *schemas.Review) error {
	return r.UpdateReview(review)
}

func (r *ReviewService) DeleteReview(id uint) error {
	return r.DeleteReview(id)
}

func (r *ReviewService) ReadReviewByBook(bookID uint) ([]schemas.Review, error) {
	return r.ReadReviewByBook(bookID)
}
