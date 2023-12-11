package model

import (
	"github.com/google/uuid"
)

type Toilet struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name       string    `json:"name"`
	Score      uint      `json:"score"`
	Location   string    `json:"location"`
	Image      string    `json:"image"`
	BuildingID uuid.UUID `gorm:"type:uuid" json:"-"`
	Building   Building  `gorm:"foreignKey:BuildingID" json:"building"`
	Reviews    []Review  `gorm:"foreignKey:ToiletID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"reviews"`
}

type CreateToiletInput struct {
	Name       string `json:"name" validate:"required"`
	Score      uint   `json:"score" validate:"required"`
	Location   string `json:"location" validate:"required"`
	Image      string `json:"image" validate:"required"`
	BuildingID string `json:"building_id" validate:"required"`
}
