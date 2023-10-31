package controller

import (
	request "district/controller/request"
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

func (cc *CartController) AddItemToCart(c echo.Context) error {
	_, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to add items to your cart.",
		})
	}

	cartToken, err := c.Cookie("cart_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing or invalid cart token.",
		})
	}

	var item request.AddCartItemRequest
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request body.",
		})
	}

	if item.ProductID == nil || item.Quantity == nil || item.Price == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing required fields.",
		})
	}

	newToken, err := cc.servicePool.CartService.AddItemToCart(cartToken.Value, item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:  "cart_token",
		Value: *newToken,
		Path:  "/",
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "item added to cart successfully.",
	})
}

func (cc *CartController) RemoveItemFromCart(c echo.Context) error {
	_, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to remove items from your cart.",
		})
	}

	cartToken, err := c.Cookie("cart_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing or invalid cart token.",
		})
	}

	var item request.RemoveCartItemRequest
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request body.",
		})
	}

	if item.ProductID == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "missing required fields.",
		})
	}

	newToken, err := cc.servicePool.CartService.RemoveItemFromCart(cartToken.Value, *item.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:  "cart_token",
		Value: *newToken,
		Path:  "/",
	})

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "item removed from cart successfully.",
	})
}
