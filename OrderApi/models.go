package models

import "time"

type Item struct {
	ItemId      uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `json:"item_code" gorm:"not null;type:varchar(8192)"`
	Description string
	Quantity    uint `gorm:"not null"`
	OrderId     uint `json:"order_id" gorm:"not null"`
}

type Order struct {
	OrderId      uint   `json:"order_id"`
	CustomerName string `json:"customer_name"`
	Items        []Item
	OrderedAt    time.Time `json:"ordered_at"`
}
