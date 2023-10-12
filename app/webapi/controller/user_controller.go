package controller

import (
	"net/http"
	"strconv"

	controller "district/controller/request"
	models "district/model"
	model "district/model/dto"
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

// - Retrieves information about the specified user (email, username, address, etc).
func (uc *UserController) GetUserInformation(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := uc.userService.GetUserByIdentification(identification)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, model.ConvertToUserDTO(user))
}

// Endpoint: PUT /api/user/:id
// - Updates the specified user with the provided information.
func (uc *UserController) UpdateUser(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := uc.userService.GetUserByIdentification(identification)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	var request controller.UpdateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user information")
	}

	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if request.Address != "" {
		user.Address = request.Address
	}
	if request.IsAdmin != nil {
		user.IsAdmin = request.IsAdmin
	}

	if err := uc.userService.UpdateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
