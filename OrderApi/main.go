package main

import (
	"fmt"
	"log"

	"excercise.id/orderapi/database"
	_ "excercise.id/orderapi/models"
)

func main() {
	database.StartDB()
	// CREATE
	// items := []models.Item{
	// 	{
	// 		ItemCode: "100216",
	// 		Quantity: 3,
	// 	},
	// 	{
	// 		ItemCode:    "106312",
	// 		Description: "using promo",
	// 		Quantity:    4,
	// 	},
	// }
	// err := database.CreateOrder("Irfan", items)
	// if err != nil {
	// 	log.Println(err)
	// }
	// GET WHERE ID
	// order, err := database.GetOrderById(4)
	// if err == nil {
	// 	log.Printf("Order data: %+v\n", order)
	// } else {
	// 	log.Println(err)
	// }
	// GET WHERE ID IN
	orders, err := database.GetOrderByIds([]uint{5, 6}...)
	if err == nil {
		log.Printf("Orders data:\n")
		for _, order := range orders {
			fmt.Printf("%+v\n", order)
		}
	} else {
		log.Println(err)
	}
}
