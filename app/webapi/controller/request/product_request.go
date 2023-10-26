package controller

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Stock       *int     `json:"stock"`
	Price       *float64 `json:"price"`
}
