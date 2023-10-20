package model

import "district/model"

type ReviewDTO struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	AuthorID  string `json:"author"`
	Content   string `json:"content"`
}

func ConvertToReviewDTO(review *model.Review) *ReviewDTO {
	return &ReviewDTO{
		ID:        review.ID,
		ProductID: review.ProductID,
		AuthorID:  review.AuthorID,
		Content:   review.Content,
	}
}
