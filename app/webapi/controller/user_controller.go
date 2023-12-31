package controller

import (
	"net/http"
	"strconv"

	request "district/controller/request"
	service "district/service/handler"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	servicePool *service.ServicePool
}

func NewUserController(servicePool *service.ServicePool) *UserController {
	return &UserController{servicePool: servicePool}
}

// Endpoint: GET /api/user
// - Retrieves information about all users (identification, email, balance, isAdmin).
func (uc *UserController) GetAllUsers(c echo.Context) error {
	token, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to check this information",
		})
	}

	users, err := uc.servicePool.UserService.GetAllUsers(token.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

// Endpoint: POST /api/user
// - Creates a new user with the specified information.
func (uc *UserController) CreateUser(c echo.Context) error {
	var request request.CreateUserRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid user information"})
	}

	user, err := uc.servicePool.UserService.CreateUser(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"user": user})
}

// Endpoint: GET /api/user/:id
// - Retrieves information about the specified user (email, username, address, etc).
func (uc *UserController) GetUserById(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid user ID"})
	}

	token, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to check this information",
		})
	}

	user, err := uc.servicePool.UserService.GetUserByIdentification(token.Value, identification)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user": user})
}

// Endpoint: PUT /api/user/:id
// - Updates the specified user with the provided information.
func (uc *UserController) UpdateUser(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid user ID"})
	}

	token, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to update an user",
		})
	}

	var request request.UpdateUserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid user information"})
	}

	if err := uc.servicePool.UserService.UpdateUser(token.Value, identification, &request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User successfully updated."})
}

// Endpoint: DELETE /api/user/:id
// - Deletes the user with given identification.
func (uc *UserController) DeleteUser(c echo.Context) error {
	identification, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid user ID"})
	}

	token, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "you must be logged in to update an user",
		})
	}

	if err := uc.servicePool.UserService.DeleteUserByIdentification(token.Value, identification); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User successfully deleted."})
}
