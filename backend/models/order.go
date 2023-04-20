package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;not null"`
	OrderDate time.Time `gorm:"not null"`
	LastOrder time.Time `gorm:"not null"`
	active    bool      `gorm:"not null; default:true"`
}

type OrderItem struct {
	gorm.Model
	ID          int     `gorm:"primaryKey;not null"`
	Name        string  `gorm:"not null;size:255"`
	OrderNumber string  `gorm:"not null; size:255"`
	Costs       float32 `gorm:"not null"`
	Payed       float32 `gorm:""`
	Comment     string  `gorm:"size:255"`
}

type Restaurant struct {
	gorm.Model
	ID          int    `gorm:"primaryKey;not null"`
	Name        string `gorm:"not null; size:255"`
	PhoneNumber string `gorm:"not null;size:255"`
	Card        string `gorm:"not null;size:255"`
}
