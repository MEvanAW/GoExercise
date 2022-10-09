package models

import "time"

type Item struct {
	ItemId      uint   `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string
	Quantity    uint
	OrderId     uint `json:"order_id"`
}

type Order struct {
	OrderId      uint      `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
}
