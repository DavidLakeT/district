package controller

import (
	request "district/controller/request"
	service "district/service/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	servicePool *service.ServicePool
}

func NewAuthController(servicePool *service.ServicePool) *AuthController {
	return &AuthController{servicePool: servicePool}
}

// Endpoint: POST /api/auth/login
// Logs in the user into the application. Retrieves token.
func (ac *AuthController) LoginUser(c echo.Context) error {
	var request request.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid login information",
		})
	}

	token, err := ac.servicePool.AuthService.Login(request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.SetCookie(&http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	})

	return c.JSON(http.StatusFound, map[string]interface{}{
		"auth_token": token,
	})
}
