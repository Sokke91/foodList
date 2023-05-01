package controllers

import (
	"foodList/models"
	"foodList/models/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrderItem(c *gin.Context) {
	var input dto.CreateOrderItemDto

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	orderItem := models.OrderItem{
		Name:        input.Name,
		OrderNumber: input.OrderNumber,
		Costs:       input.Costs,
		Payed:       input.Payed,
		Comment:     input.Comment,
	}

	saveOrderItem, err := orderItem.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": saveOrderItem})
}

func GetOrderItemById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	orderItem, err := models.GetOrderItemById(idInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": orderItem})
	}
}
