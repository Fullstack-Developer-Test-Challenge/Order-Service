package repository

import (
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetByProductID(productID int) ([]models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetByProductID(productId int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("product_id = ?", productId).Find(&orders).Error
	return orders, err
}
