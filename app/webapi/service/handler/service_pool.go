package service

import "district/service"

type ServicePool struct {
	AuthService    *service.AuthService
	CartService    *service.CartService
	ProductService *service.ProductService
	ReviewService  *service.ReviewService
	UserService    *service.UserService
	UtilsService   *service.UtilsService
}

func NewServicePool(
	authService *service.AuthService,
	cartService *service.CartService,
	productService *service.ProductService,
	reviewService *service.ReviewService,
	userService *service.UserService,
	utilsService *service.UtilsService,
) *ServicePool {
	return &ServicePool{
		AuthService:    authService,
		CartService:    cartService,
		ProductService: productService,
		ReviewService:  reviewService,
		UserService:    userService,
		UtilsService:   utilsService,
	}
}
