package models

import (
	"foodList/database"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null; size:255"`
	PhoneNumber string `gorm:"not null;size:255"`
	Card        string `gorm:"not null;size:255"`
}

func (r *Restaurant) Save() (*Restaurant, error) {
	err := database.DB.Create(&r).Error
	if err != nil {
		return &Restaurant{}, err
	}
	return r, nil
}

func GetRestaurantById(id int) (Restaurant, error) {
	var restaurant Restaurant
	err := database.DB.First(&restaurant, id)
	if err != nil {
		return Restaurant{}, nil
	}
	return restaurant, nil
}

func (r *Restaurant) Delete() error {
	err := database.DB.Delete(&r).Error
	if err != nil {
		return err
	}
	return nil
}
