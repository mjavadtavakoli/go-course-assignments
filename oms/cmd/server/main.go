package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	httpAdapter "oms/internal/adapters/http"
	"oms/internal/adapters/repository"
	"oms/internal/core/services"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/oms_db?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("database ping failed:", err)
	}

	orderRepo := repository.NewMySQLOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := httpAdapter.NewOrderHandler(orderService)

	router := gin.Default()

	oms := router.Group("/oms/api")
	{
		oms.POST("/orders", orderHandler.CreateOrder)
		oms.GET("/orders/:id", orderHandler.GetOrderByID)
	}

	log.Println("server is running on :8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("failed to run server:", err)
	}
}
