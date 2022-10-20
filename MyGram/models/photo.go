package models

type Photo struct {
	Model
	Title    string `gorm:"not null" validate:"required"`
	Caption  string
	PhotoUrl string `gorm:"not null" validate:"required"`
	UserID   uint
}
