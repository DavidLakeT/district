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

func (cc *CartController) GetCartInformation(c echo.Context) error {
	_, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to see your cart.",
		})
	}

	cartToken, err := c.Cookie("cart_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing or invalid cart token.",
		})
	}

	totalPrice, cartInformation, err := cc.servicePool.CartService.GetCartInformation(cartToken.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"cart":        cartInformation,
	})
}
