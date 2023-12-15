package model

import (
	"github.com/google/uuid"
)

type Review struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username string    `json:"username"`
	Comment  string    `json:"comment"`
	Score    uint      `json:"score"`
	ToiletID uuid.UUID `gorm:"type:uuid" json:"-"`
}

type CreateReviewInput struct {
	Username string `json:"username"`
	Comment  string `json:"comment"`
	Score    uint   `json:"score"`
	ToiletID string `json:"toiletId"`
}
