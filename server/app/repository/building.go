package repository

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/platform/database"
)

type BuildingRepository struct{}

func NewBuildingRepository() BuildingRepository {
	return BuildingRepository{}
}

func (r *BuildingRepository) FindAll() ([]model.Building, error) {
	db := database.GetDB()

	var buildings []model.Building

	if err := db.Find(&buildings).Error; err != nil {
		return nil, err
	}

	return buildings, nil
}

func (r *BuildingRepository) FindByID(id string) (model.Building, error) {
	db := database.GetDB()

	var building model.Building

	if err := db.First(&building, id).Error; err != nil {
		return model.Building{}, err
	}

	return building, nil
}

func (r *BuildingRepository) Create(building model.Building) (model.Building, error) {
	db := database.GetDB()

	if err := db.Create(&building).Error; err != nil {
		return model.Building{}, err
	}

	return building, nil
}
