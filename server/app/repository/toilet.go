package repository

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/platform/database"
)

type ToiletRepository struct{}

func NewToiletRepository() ToiletRepository {
	return ToiletRepository{}
}

func (r *ToiletRepository) FindAll() ([]model.Toilet, error) {
	db := database.GetDB()

	var toilets []model.Toilet

	if err := db.Find(&toilets).Error; err != nil {
		return nil, err
	}

	return toilets, nil
}

func (r *ToiletRepository) FindByID(id string) (model.Toilet, error) {
	db := database.GetDB()

	var toilet model.Toilet

	if err := db.First(&toilet, id).Error; err != nil {
		return model.Toilet{}, err
	}

	return toilet, nil
}

func (r *ToiletRepository) Create(toilet model.Toilet) (model.Toilet, error) {
	db := database.GetDB()

	if err := db.Create(&toilet).Error; err != nil {
		return model.Toilet{}, err
	}

	return toilet, nil
}
