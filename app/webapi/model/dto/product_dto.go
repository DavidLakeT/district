package model

import "district/model"

type ProductDTO struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	Reviews     []*model.Review `json:"reviews"`
}

func ConvertToProductDTO(product *model.Product) *ProductDTO {
	return &ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Reviews:     product.Reviews,
	}
}
