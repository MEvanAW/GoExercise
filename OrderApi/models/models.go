package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"not null;type:varchar(8192)" json:"item_code"`
	Description string `gorm:"type:varchar(8192)"`
	Quantity    uint   `gorm:"not null"`
	OrderID     uint   `json:"order_id"`
}

type Order struct {
	ID           uint   `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"type:varchar(8192)" json:"customer_name"`
	Items        []Item
	OrderedAt    time.Time `gorm:"not null" json:"ordered_at"`
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
