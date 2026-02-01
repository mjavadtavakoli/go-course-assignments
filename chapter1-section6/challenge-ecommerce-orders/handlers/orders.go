package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ecommerce-orders/models"
	"ecommerce-orders/utils"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateOrder(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.ID = uuid.NewString()
	order.Status = "pending"
	orders[order.ID] = order

	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	result := []models.Order{}
	for _, o := range orders {
		result = append(result, o)
	}
	c.JSON(http.StatusOK, result)
}
