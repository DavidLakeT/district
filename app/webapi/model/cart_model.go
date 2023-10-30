package model

type CartItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Cost      float64 `json:"cost"`
}

type Cart []CartItem
