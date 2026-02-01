package main

import (
	"github.com/gin-gonic/gin"
	"ecommerce-orders/handlers"
)

func main() {
	r := gin.Default()

	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)

	r.Run(":8080")
}
