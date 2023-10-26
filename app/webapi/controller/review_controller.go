package controller

import (
	request "district/controller/request"
	"district/model"
	service "district/service/handler"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	servicePool *service.ServicePool
}

func NewReviewController(servicePool *service.ServicePool) *ReviewController {
	return &ReviewController{servicePool: servicePool}
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

	token, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to create a review",
		})
	}

	if err := rc.servicePool.ReviewService.CreateReview(token.Value, &review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"product_id": review.ProductID,
		"content":    review.Content,
	})
}

// Endpoint: GET /api/review/:id
// - Retrieves the review with the specified ID.
func (rc *ReviewController) GetReviewById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid product ID",
		})
	}

	review, err := rc.servicePool.ReviewService.GetReviewById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "review not found",
		})
	}

	return c.JSON(http.StatusOK, review)
}

// Endpoint: DELETE /api/review/:id
// - Deletes the review with the specified ID.
func (rc *ReviewController) DeleteReviewById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid review ID",
		})
	}

	if err := rc.servicePool.ReviewService.DeleteReview(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "review succesfully deleted."})
}
