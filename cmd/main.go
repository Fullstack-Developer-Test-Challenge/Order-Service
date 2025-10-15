package main

import (
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/cache"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/config"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/handler"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/models"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/repository"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/routes"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	_ = godotenv.Load()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.Order{})

	cache := cache.NewRedisCache("localhost:6379")

	orderRepo := repository.NewOrderRepository(config.DB)
	orderService := service.NewOrderService(orderRepo, cache)
	orderHandler := handler.NewOrderHandler(orderService)

	r := gin.Default()
	routes.OrderRoutes(r, orderHandler)

	r.Run()
}
