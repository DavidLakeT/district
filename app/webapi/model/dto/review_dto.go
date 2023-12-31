package model

import "district/model"

type ReviewDTO struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	UserID    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	Content   string `json:"content"`
}

func ConvertToReviewDTO(review *model.Review) *ReviewDTO {
	return &ReviewDTO{
		ID:        review.ID,
		ProductID: review.ProductID,
		UserID:    review.UserID,
		UserEmail: review.UserEmail,
		Content:   review.Content,
	}
}
