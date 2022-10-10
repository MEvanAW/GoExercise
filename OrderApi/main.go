package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"excercise.id/orderapi/database"
	"excercise.id/orderapi/models"
	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	// CREATE
	// items := []models.Item{
	// 	{
	// 		ItemCode: "100214",
	// 		Quantity: 1,
	// 	},
	// 	{
	// 		ItemCode:    "106310",
	// 		Description: "using promo",
	// 		Quantity:    3,
	// 	},
	// }
	// createOrder("Evan", items)
	// GET WHERE ID
	log.Printf("Order data: %+v\n", getOrderById(1))
}

func createOrder(customerName string, items []models.Item) {
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

func getOrderById(id uint) models.Order {
	db := database.GetDB()
	order := models.Order{}
	err := db.Model(&models.Order{}).Preload("Items").First(&order, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Order with id", id, "not found.")
		}
		log.Println("Error finding order:", err)
	}
	return order
}
