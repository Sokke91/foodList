package dto

import "time"

type CreateOrderDto struct {
	OrderDate time.Time `json:"orderDate"`
	LastOrder time.Time `json:"lastOrder"`
	Active    bool      `json:"active"`
}
