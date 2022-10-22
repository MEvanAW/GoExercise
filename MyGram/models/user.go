package models

type User struct {
	Model
	Username string `gorm:"not null;uniqueIndex"`
	Email    string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Age      uint   `gorm:"not null"`
}
