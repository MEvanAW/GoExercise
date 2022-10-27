package models

type Comment struct {
	Model
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
	Message string `gorm:"not null;type:varchar(8192)" json:"message"`
}
