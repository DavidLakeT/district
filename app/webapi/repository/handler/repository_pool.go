package repository

import "district/repository"

type RepositoryPool struct {
	ProductRepository *repository.ProductRepository
	ReviewRepository  *repository.ReviewRepository
	UserRepository    *repository.UserRepository
}

func NewRepositoryPool(
	productRepository *repository.ProductRepository,
	reviewRepository *repository.ReviewRepository,
	userRepository *repository.UserRepository,
) *RepositoryPool {
	return &RepositoryPool{
		ProductRepository: productRepository,
		ReviewRepository:  reviewRepository,
		UserRepository:    userRepository,
	}
}
