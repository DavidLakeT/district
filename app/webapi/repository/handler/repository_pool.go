package repository

import "district/repository"

type RepositoryPool struct {
	AuthRepository    *repository.AuthRepository
	ProductRepository *repository.ProductRepository
	ReviewRepository  *repository.ReviewRepository
	UserRepository    *repository.UserRepository
}

func NewRepositoryPool(
	authRepository *repository.AuthRepository,
	productRepository *repository.ProductRepository,
	reviewRepository *repository.ReviewRepository,
	userRepository *repository.UserRepository,
) *RepositoryPool {
	return &RepositoryPool{
		AuthRepository:    authRepository,
		ProductRepository: productRepository,
		ReviewRepository:  reviewRepository,
		UserRepository:    userRepository,
	}
}
