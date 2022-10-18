package models

type Photo struct {
	Model
	Title    string
	Caption  string
	PhotoUrl string
	UserID   uint
}
