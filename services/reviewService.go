package services

import (
	"github.com/oliverperboni/GoTomekeeper/repository"
	"github.com/oliverperboni/GoTomekeeper/schemas"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func CreateReviewService(r repository.ReviewRepository) ReviewService {
	return ReviewService{repo: r}
}

func (r *ReviewService) CreateReview(review *schemas.Review) error {
	return r.repo.CreateReview(review)
}

func (r *ReviewService) UpdateReview(review *schemas.Review) error {
	return r.repo.UpdateReview(review)
}

func (r *ReviewService) DeleteReview(id uint) error {
	return r.repo.DeleteReview(id)
}

func (r *ReviewService) ReadReviewByBook(bookID uint) ([]schemas.Review, error) {
	return r.repo.ReadReviewByBook(bookID)
}
