package models

type SocialMedia struct {
	Model
	Name           string `gorm:"type:varchar(8192)"`
	SocialMediaUrl string
	UserID         uint
}
