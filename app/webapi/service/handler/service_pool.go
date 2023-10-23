package service

import "district/service"

type ServicePool struct {
	AuthService    *service.AuthService
	ProductService *service.ProductService
	ReviewService  *service.ReviewService
	UserService    *service.UserService
}

func NewServicePool(
	authService *service.AuthService,
	productService *service.ProductService,
	reviewService *service.ReviewService,
	userService *service.UserService,
) *ServicePool {
	return &ServicePool{
		AuthService:    authService,
		ProductService: productService,
		ReviewService:  reviewService,
		UserService:    userService,
	}
}
