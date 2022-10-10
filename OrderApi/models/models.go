package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(8192)"`
	Description string `gorm:"type:varchar(8192)"`
	Quantity    uint   `gorm:"not null"`
	OrderID     uint
}

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"type:varchar(8192)"`
	Items        []Item
	OrderedAt    time.Time `gorm:"not null"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ItemCode == "" {
		err = errors.New("ItemCode may not be empty.")
	}
	return
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if o.CustomerName == "" {
		err = errors.New("CustomerName may not be empty.")
	}
	return
}
