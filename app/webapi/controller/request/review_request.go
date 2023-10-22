package controller

type CreateReviewRequest struct {
	ProductId int    `json:"product_id"`
	Content   string `json:"content"`
}
