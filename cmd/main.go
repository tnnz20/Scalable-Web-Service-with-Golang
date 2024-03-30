package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tnnz20/Scalable-Web-Service-with-Golang/config"
	"github.com/tnnz20/Scalable-Web-Service-with-Golang/internal/order"
	"github.com/tnnz20/Scalable-Web-Service-with-Golang/pkg/datasource"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("local")
	if err != nil {
		panic(err)
	}

	// Create Connection Database & Migrate
	db, err := datasource.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	db.Migrate(&order.Order{}, &order.Item{})

	// Create Gin order
	router := gin.Default()

	orderRepo := order.NewRepository(db.GetDB())
	orderSvc := order.NewService(orderRepo)
	orderHandler := order.NewHandler(orderSvc)

	// Route
	order := router.Group("/orders")

	order.POST("", orderHandler.CreateOrder)
	order.GET("", orderHandler.GetOrder)
	order.PUT("/:id", orderHandler.UpdateOrder)
	order.DELETE("/:id", orderHandler.DeleteOrder)

	router.Run(":8080")
}
