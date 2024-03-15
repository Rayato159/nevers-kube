package entities

type Image struct {
	ID          string `gorm:"primaryKey" json:"id"`
	ImageBase64 string `json:"imageBase64" validate:"required"`
}
