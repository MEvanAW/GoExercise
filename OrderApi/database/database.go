package database

import (
	"errors"
	"fmt"
	"log"
	"time"

	"excercise.id/orderapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "postgres"
	dbPort = "5432"
	dbName = "order-api"
	db     *gorm.DB
	err    error
)

func StartDB() {
	var password string
	fmt.Print("Enter db password: ")
	fmt.Scanln(&password)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}

func CreateOrder(customerName string, items []models.Item) error {
	if db == nil {
		return errors.New("DB hasn't started yet.")
	}
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

func GetOrderById(id uint) (models.Order, error) {
	order := models.Order{}
	if db == nil {
		return order, errors.New("DB hasn't started yet.")
	}
	err := db.Model(&models.Order{}).Preload("Items").Take(&order, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, errors.New(fmt.Sprintf("Order with ID %d not found.", id))
		}
		return order, err
	}
	return order, nil
}

func GetOrderByIds(ids ...uint) ([]models.Order, error) {
	if len(ids) == 1 {
		order, err := GetOrderById(ids[0])
		return []models.Order{order}, err
	}
	if db == nil {
		return nil, errors.New("DB hasn't started yet.")
	}
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
