package repository

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/platform/database"
)

type ReviewRepository struct{}

func NewReviewRepository() ReviewRepository {
	return ReviewRepository{}
}

func (r *ReviewRepository) FindAll() ([]model.Review, error) {
	db := database.GetDB()

	var reviews []model.Review

	if err := db.Find(&reviews).Error; err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) FindByID(id string) (model.Review, error) {
	db := database.GetDB()

	var review model.Review

	if err := db.First(&review, id).Error; err != nil {
		return model.Review{}, err
	}

	return review, nil
}

func (r *ReviewRepository) Create(review model.Review) (model.Review, error) {
	db := database.GetDB()

	if err := db.Create(&review).Error; err != nil {
		return model.Review{}, err
	}

	return review, nil
}
