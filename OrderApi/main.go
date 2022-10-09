package main

import (
	"fmt"
	"log"
	"time"

	"excercise.id/orderapi/database"
	"excercise.id/orderapi/models"
)

func main() {
	database.StartDB()
	items := []models.Item{
		{
			ItemCode: "100214",
			Quantity: 1,
		},
		{
			ItemCode:    "106310",
			Description: "using promo",
			Quantity:    3,
		},
	}
	CreateOrder("Evan", items)
}

func CreateOrder(customerName string, items []models.Item) {
	db := database.GetDB()
	Order := models.Order{
		CustomerName: customerName,
		Items:        items,
		OrderedAt:    time.Now(),
	}
	err := db.Create(&Order).Error
	if err != nil {
		fmt.Println("Error creating order data: ", err)
	}
	log.Println("New Order Data: ", Order)
}
