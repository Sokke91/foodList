package dto

type CreateOrderItemDto struct {
	Name        string  `json:"name"`
	OrderNumber string  `json:"ordernumber"`
	Costs       float32 `json:"costs"`
	Payed       float32 `json:"payed"`
	Comment     string  `json:"comment"`
}
