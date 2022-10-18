package models

type Comment struct {
	Model
	UserID  uint
	PhotoID uint
	Message string
}
