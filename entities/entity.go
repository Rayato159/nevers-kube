package entities

import "github.com/google/uuid"

type Image struct {
	ID          uuid.UUID `gorm:"primaryKey;autoIncrement;type:uuid" json:"id"`
	ImageBase64 string    `json:"iamgeBase64" validate:"required"`
}
