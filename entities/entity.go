package entities

type Image struct {
	ID          string `gorm:"primaryKey;autoIncrement" json:"id"`
	ImageBase64 string `json:"iamgeBase64" validate:"required"`
}
