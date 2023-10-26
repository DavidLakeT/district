package model

import (
	"district/model"
)

type ProductDTO struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Stock       int          `json:"stock"`
	Price       float64      `json:"price"`
	Reviews     []*ReviewDTO `json:"reviews"`
}

func ConvertToProductDTO(product *model.Product) *ProductDTO {
	reviewDtos := make([]*ReviewDTO, len(product.Reviews))
	for i, review := range product.Reviews {
		reviewDtos[i] = ConvertToReviewDTO(review)
	}

	return &ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
		Price:       product.Price,
		Reviews:     reviewDtos,
	}
}
