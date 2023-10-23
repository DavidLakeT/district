package controller

import "district/controller"

type ControllerPool struct {
	AuthController    *controller.AuthController
	ProductController *controller.ProductController
	ReviewController  *controller.ReviewController
	UserController    *controller.UserController
}

func NewControllerPool(
	authController *controller.AuthController,
	productController *controller.ProductController,
	reviewController *controller.ReviewController,
	userController *controller.UserController,
) *ControllerPool {
	return &ControllerPool{
		AuthController:    authController,
		ProductController: productController,
		ReviewController:  reviewController,
		UserController:    userController,
	}
}
