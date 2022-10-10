package main

import (
	_ "fmt"
	"log"
	_ "time"

	"excercise.id/orderapi/database"
	_ "excercise.id/orderapi/models"
)

func main() {
	database.StartDB()
	// CREATE
	// items := []models.Item{
	// 	{
	// 		ItemCode: "100217",
	// 		Quantity: 4,
	// 	},
	// 	{
	// 		ItemCode:    "106313",
	// 		Description: "using promo",
	// 		Quantity:    5,
	// 	},
	// }
	// err := database.CreateOrder("Tyo", items)
	// if err != nil {
	// 	log.Println(err)
	// }
	// GET WHERE ID
	// order, err := database.GetOrderById(5)
	// if err == nil {
	// 	log.Printf("Order data: %+v\n", order)
	// } else {
	// 	log.Println(err)
	// }
	// GET WHERE ID IN
	// orders, err := database.GetOrderByIds([]uint{5, 6}...)
	// if err == nil {
	// 	log.Printf("Orders data:\n")
	// 	for _, order := range orders {
	// 		fmt.Printf("%+v\n", order)
	// 	}
	// } else {
	// 	log.Println(err)
	// }
	// UPDATE
	// customerName := "Tyo"
	// items := []models.Item{
	// 	{
	// 		ItemCode: "UPDATEDCODE",
	// 		Quantity: 2,
	// 	},
	// }
	// orderedAt := time.Now()
	// err := database.UpdateOrderById(5, customerName, items, orderedAt)
	// if err != nil {
	// 	log.Println(err)
	// }
	// DELETE
	err := database.DeleteOrderById(5)
	if err != nil {
		log.Println(err)
	}
}
