package models

import (
	"foodList/database"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;not null"`
	OrderDate time.Time `gorm:"not null"`
	LastOrder time.Time `gorm:"not null"`
	Active    bool      `gorm:"not null; default:true"`
}

// Order Model
func (order *Order) Save() (*Order, error) {
	err := database.DB.Create(&order).Error
	if err != nil {
		return &Order{}, err
	}
	return order, nil
}

func GetOrderById(id int) (Order, error) {
	var order Order
	err := database.DB.First(&order, id).Error
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func (order *Order) Delete() error {
	err := database.DB.Delete(&order).Error
	if err != nil {
		return err
	}
	return nil
}
