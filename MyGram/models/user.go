package models

type User struct {
	Model
	Username string
	Email    string
	Password string
	Age      uint
}
