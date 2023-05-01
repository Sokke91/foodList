package dto

type CreateRestaurantDto struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Card        string `json:"card"`
}
