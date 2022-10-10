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
	// 		ItemCode: "100215",
	// 		Quantity: 2,
	// 	},
	// 	{
	// 		ItemCode:    "106311",
	// 		Description: "on discount",
	// 		Quantity:    4,
	// 	},
	// }
	// err := createOrder("Fandi", items)
	// if err != nil {
	// 	log.Println(err)
	// }
	// GET WHERE ID
	// order, err := getOrderById(3)
	// if err == nil {
	// 	log.Printf("Order data: %+v\n", order)
	// } else {
	// 	log.Println(err)
	// }
	// GET WHERE ID IN
	orders, err := getOrderByIds([]uint{4, 5}...)
	if err == nil {
		log.Printf("Orders data:\n")
		for _, order := range orders {
			fmt.Printf("%+v\n", order)
		}
	} else {
		log.Println(err)
	}
}

func createOrder(customerName string, items []models.Item) error {
	db := database.GetDB()
	Order := models.Order{
		CustomerName: customerName,
		Items:        items,
		OrderedAt:    time.Now(),
	}
	err := db.Create(&Order).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Error creating order data: %s", err))
	}
	log.Println("New Order Data: ", Order)
	return nil
}

func getOrderById(id uint) (models.Order, error) {
	db := database.GetDB()
	order := models.Order{}
	err := db.Model(&models.Order{}).Preload("Items").Take(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, errors.New(fmt.Sprintf("Order with ID %d not found.", id))
		}
		return order, err
	}
	return order, nil
}

func getOrderByIds(ids ...uint) ([]models.Order, error) {
	if len(ids) == 1 {
		order, err := getOrderById(ids[0])
		return []models.Order{order}, err
	}
	db := database.GetDB()
	orders := make([]models.Order, 2)
	err := db.Model(&models.Order{}).Preload("Items").Find(&orders, ids).Error
	if err != nil {
		return nil, err
	}
	if len(orders) == 0 {
		return nil, errors.New("No ID match.")
	}
	return orders, nil
}
