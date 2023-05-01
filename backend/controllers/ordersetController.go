package controllers

import (
	"foodList/models"
	"foodList/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context, id int) {
	order, err := models.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func CreateOrderset(c *gin.Context) {
	var userInput dto.CreateOrderDto

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	order := models.Order{OrderDate: userInput.OrderDate, LastOrder: userInput.LastOrder, Active: true}
	saveOrder, err := order.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": saveOrder})
}

func DeleteOrderset(c *gin.Context) {
}
