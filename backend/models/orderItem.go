package models

import (
	"foodList/database"

	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ID          int     `gorm:"primaryKey;not null"`
	Name        string  `gorm:"not null;size:255"`
	OrderNumber string  `gorm:"not null; size:255"`
	Costs       float32 `gorm:"not null"`
	Payed       float32 `gorm:""`
	Comment     string  `gorm:"size:255"`
}

func (order *OrderItem) Save() (*OrderItem, error) {
	err := database.DB.Create(&order).Error
	if err != nil {
		return &OrderItem{}, err
	}
	return order, nil
}

func GetOrderItemById(id int) (OrderItem, error) {
	var orderItem OrderItem
	err := database.DB.First(&orderItem, id).Error
	if err != nil {
		return OrderItem{}, err
	}
	return orderItem, nil
}

func (orderItem *OrderItem) Delete() error {
	err := database.DB.Delete(&orderItem).Error
	if err != nil {
		return err
	}
	return nil
}
