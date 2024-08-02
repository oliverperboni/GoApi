package repository

import (
	"github.com/oliverperboni/GoTomekeeper/schemas"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	DB *gorm.DB
}

func CreteNewReviewRepo(db *gorm.DB) ReviewRepository {
	return ReviewRepository{DB: db}
}

func (r *ReviewRepository) CreateReview(review *schemas.Review) error {
	return r.DB.Create(review).Error
}

func (r *ReviewRepository) UpdateReview(review *schemas.Review) error {
	return r.DB.Save(review).Error
}

func (r *ReviewRepository) DeleteReview(id uint) error {
	return r.DB.Delete("Id = ?", id).Error
}

func (r *ReviewRepository) ReadReviewByBook(bookID uint) ([]schemas.Review, error) {
	var listReview []schemas.Review
	err := r.DB.Find(&listReview, "book_id + ?", bookID).Error
	return listReview, err
}
