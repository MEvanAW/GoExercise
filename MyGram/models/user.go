package models

type User struct {
	Model
	Username string `gorm:"not null;uniqueIndex" validate:"required"`
	Email    string `gorm:"not null;uniqueIndex" validate:"required,email"`
	Password string `gorm:"not null" validate:"required,min=6"`
	Age      uint   `gorm:"not null" validate:"required,gt=8"`
}
