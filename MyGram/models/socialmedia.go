package models

type SocialMedia struct {
	Model
	Name           string `gorm:"not null;type:varchar(8192)" validate:"required"`
	SocialMediaUrl string `gorm:"not null;type:varchar(8192)" validate:"required"`
	UserID         uint
}
