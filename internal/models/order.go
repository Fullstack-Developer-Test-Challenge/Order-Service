package models

import "time"

type Order struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	ProductID  int       `json:"product_id"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
