package controller

import (
	service "district/service/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	servicePool *service.ServicePool
}

func NewCartController(servicePool *service.ServicePool) *CartController {
	return &CartController{servicePool: servicePool}
}

func (cc *CartController) GetCartTotal(c echo.Context) error {
	_, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to get your cart total",
		})
	}

	cartToken, err := c.Cookie("cart_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "could not read your cart information",
		})
	}

	cartTotal, err := cc.servicePool.CartService.GetCartTotal(cartToken.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": cartTotal,
	})
}