package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/mjavadtavakoli/challenge-ecommerce-orders/models"	
)

var orders []models.Order

func GetOrders(c *gin.Context) {
    c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    orders = append(orders, order)
    c.JSON(http.StatusCreated, order)
}