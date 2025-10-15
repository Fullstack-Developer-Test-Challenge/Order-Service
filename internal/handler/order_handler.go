package handler

import (
	"fmt"

	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/dto"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order, err := h.service.CreateOrder(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, order)
}

func (h *OrderHandler) GetByProductID(c *gin.Context) {
	productIDParam := c.Query("product_id")
	if productIDParam == "" {
		c.JSON(400, gin.H{"error": "product_id parameter is required"})
		return
	}

	var productID int
	if _, err := fmt.Sscanf(productIDParam, "%d", &productID); err != nil {
		c.JSON(400, gin.H{"error": "invalid product_id parameter"})
		return
	}

	orders, err := h.service.GetByProductID(productID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, orders)
}
