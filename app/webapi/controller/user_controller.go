package controller

import (
	"net/http"
	"strconv"

	controller "district/controller/request"
	"district/models"
	"district/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Endpoint: GET /api/user/:id
// - Retrieves information about the specified user (email, username, address, etc).
func (uc *UserController) UserInformation(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := uc.userService.GetUserByIdentification(identification)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// Endpoint: POST /api/user
// - Creates a new user with the specified information.
func (uc *UserController) CreateUser(c echo.Context) error {
	var request controller.CreateUserRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user information")
	}

	user := models.User{
		Identification: request.Identification,
		Email:          request.Email,
		Username:       request.Username,
		Password:       request.Password,
		Address:        request.Address,
		Balance:        0,
		IsAdmin:        request.IsAdmin,
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}
