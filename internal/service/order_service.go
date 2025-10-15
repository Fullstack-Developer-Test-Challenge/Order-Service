package service

import (
	"log"

	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/cache"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/client"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/dto"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/models"
	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/repository"
)

type OrderService interface {
	CreateOrder(order dto.CreateOrderRequest) (dto.CreateOrderResponse, error)
	GetByProductID(productID int) ([]models.Order, error)
}

type orderService struct {
	repo  repository.OrderRepository
	cache *cache.RedisCache
}

func NewOrderService(repo repository.OrderRepository, cache *cache.RedisCache) OrderService {
	return &orderService{repo, cache}
}

func (s *orderService) CreateOrder(req dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
	// validasi product
	_, err := client.GetByProductID(req.ProductID)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}

	order := models.Order{
		ProductID:  req.ProductID,
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
	}

	if err := s.repo.CreateOrder(&order); err != nil {
		return dto.CreateOrderResponse{}, err
	}
	return dto.CreateOrderResponse{
		ID:         order.ID,
		ProductID:  order.ID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}, nil
}

func (s *orderService) GetByProductID(productID int) ([]models.Order, error) {
	// cek cache
	cached, err := s.cache.GetOrders(productID)
	if err == nil && cached != nil {
		log.Print("ambil dari cache")
		return cached, nil
	}

	orders, err := s.repo.GetByProductID(productID)
	if err != nil {
		return nil, err
	}

	_ = s.cache.SetOrders(productID, orders)

	log.Print("ambil baru")
	return orders, nil
}
