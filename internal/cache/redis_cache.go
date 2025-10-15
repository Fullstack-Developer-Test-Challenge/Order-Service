package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Fullstack-Developer-Test-Challenge/Order-Service/internal/models"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(addr string) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) GetOrders(productId int) ([]models.Order, error) {
	key := fmt.Sprintf("orders:products:%d", productId)
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var orders []models.Order
	if err := json.Unmarshal([]byte(val), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *RedisCache) SetOrders(productId int, orders []models.Order) error {
	key := fmt.Sprintf("orders:products:%d", productId)
	data, _ := json.Marshal(orders)
	return r.client.Set(r.ctx, key, data, 5*time.Minute).Err()
}
