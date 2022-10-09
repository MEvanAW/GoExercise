package models

import "time"

type Item struct {
	ItemId      uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"not null;type:varchar(8192)" json:"item_code"`
	Description string `gorm:"type:varchar(8192)"`
	Quantity    uint   `gorm:"not null"`
	OrderId     uint   `gorm:"not null" json:"order_id"`
}

type Order struct {
	OrderId      uint   `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"type:varchar(8192)" json:"customer_name"`
	Items        []Item
	OrderedAt    time.Time `gorm:"not null" json:"ordered_at"`
}
