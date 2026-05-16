package main

import (
	httpHandler "oms/internal/adapters/http"
	"oms/internal/adapters/repository"
	"oms/internal/database"
	"oms/internal/logger"
	"oms/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.NewMySQLConnection()

	repo := repository.NewMySQLOrderRepository(db)

	orderService := service.NewOrderService(repo)

	handler := httpHandler.NewHandler(orderService)

	router := gin.Default()

	//logger
	err := logger.InitLogger()
	if err != nil {
		panic(err)
	}
	logger.Log.Println("application started")

	router.POST("/oms/api/orders", handler.CreateOrder)
	router.GET("/oms/api/orders/:id", handler.GetOrder)
	router.GET("/oms/api/orders", handler.GetOrders)
	router.PUT("/oms/api/orders/:id", handler.UpdateOrder)

	router.Run(":8080")

}
