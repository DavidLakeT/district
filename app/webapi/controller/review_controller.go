package controller

import (
	request "district/controller/request"
	"district/model"
	"district/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	reviewService *service.ReviewService
}

func NewReviewController(reviewService *service.ReviewService) *ReviewController {
	return &ReviewController{reviewService: reviewService}
}

// Endpoint: POST /api/review
// - Creates a new product review with the specified information.
func (rc *ReviewController) CreateReview(c echo.Context) error {
	var request request.CreateReviewRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid review information",
		})
	}

	review := model.Review{
		ProductID: request.ProductId,
		Content:   request.Content,
	}

	if err := rc.reviewService.CreateReview(&review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"product_id": review.ProductID,
		"content":    review.Content,
	})
}
