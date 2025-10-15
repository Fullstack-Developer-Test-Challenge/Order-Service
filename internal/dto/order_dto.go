package dto

type CreateOrderRequest struct {
	ProductID  int     `json:"product_id" binding:"required"`
	TotalPrice float64 `json:"total_price" binding:"required,gt=0"`
	Status     string  `json:"status" binding:"required,oneof=pending paid canceled"`
}

type CreateOrderResponse struct {
	ID         int     `json:"id"`
	ProductID  int     `json:"product_id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
