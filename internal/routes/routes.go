package routes

import (
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/handler"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine, orderHandler *handler.OrderHandler) {
	orders := r.Group("/orders")
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("", orderHandler.GetByProductID)
	}
}
