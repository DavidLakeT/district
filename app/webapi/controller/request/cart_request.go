package controller

type AddCartItemRequest struct {
	ProductID *int     `json:"product_id"`
	Quantity  *int     `json:"quantity"`
	Price     *float64 `json:"price"`
}

type RemoveCartItemRequest struct {
	ProductID *int `json:"product_id"`
}
