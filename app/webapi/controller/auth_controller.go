package controller

import (
	controller "district/controller/request"
	"district/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Endpoint: POST /api/auth/login
// Logs in the user into the application. Retrieves token.
func (ac *AuthController) LoginUser(c echo.Context) error {
	var request controller.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid login information")
	}

	token, err := ac.authService.Login(request.Email, request.Password)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, token)
}
