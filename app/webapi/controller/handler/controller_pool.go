package controller

import "district/controller"

type ControllerPool struct {
	AuthController    *controller.AuthController
	CartController    *controller.CartController
	ProductController *controller.ProductController
	ReviewController  *controller.ReviewController
	UserController    *controller.UserController
	UtilsController   *controller.UtilsController
}

func NewControllerPool(
	authController *controller.AuthController,
	cartController *controller.CartController,
	productController *controller.ProductController,
	reviewController *controller.ReviewController,
	userController *controller.UserController,
	utilsController *controller.UtilsController,
) *ControllerPool {
	return &ControllerPool{
		AuthController:    authController,
		CartController:    cartController,
		ProductController: productController,
		ReviewController:  reviewController,
		UserController:    userController,
		UtilsController:   utilsController,
	}
}
