package controller

import (
	"fmt"
	"net/http"
	"strconv"

	controller "district/controller/request"
	models "district/model"
	model "district/model/dto"
	service "district/service"

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
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "Invalid user information")
	}

	hashedPassword, err := service.HashPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := models.User{
		Identification: request.Identification,
		Email:          request.Email,
		Username:       request.Username,
		Password:       hashedPassword,
		Address:        request.Address,
		Balance:        0,
		IsAdmin:        request.IsAdmin,
	}

	if err := uc.userService.CreateUser(&user); err != nil {
		fmt.Println("error:", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, model.ConvertToUserDTO(&user))
}

// Endpoint: GET /api/user/:id
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

	return c.JSON(http.StatusOK, user)
}

// Endpoint: PUT /api/user/:id
// - Updates the specified user with the provided information.
func (uc *UserController) UpdateUser(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var request controller.UpdateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user information")
	}

	if err := uc.userService.UpdateUser(identification, &request); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product succesfully updated.",
	})
}

// Endpoint: DELETE /api/user/:id
// - Deletes the user with given identification.
func (uc *UserController) DeleteUser(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	err = uc.userService.DeleteUserByIdentification(identification)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}
