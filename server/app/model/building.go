package model

import (
	"github.com/google/uuid"
)

type Building struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name    string    `json:"name"`
	Toilets []Toilet  `gorm:"foreignKey:BuildingID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"toilets"`
}

type CreateBuildingInput struct {
	Name string `json:"name" validate:"required"`
}
