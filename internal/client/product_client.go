package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID        int
	Name      string
	Price     float64
	Qty       int
	CreatedAt string
}

func GetByProductID(productId int) (*Product, error) {
	url := fmt.Sprintf("http://localhost:3000/products/%d", productId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product not found (status: %d)", resp.StatusCode)
	}

	var product Product
	if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
		return nil, err
	}
	return &product, err
}
