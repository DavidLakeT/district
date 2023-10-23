package repository

import "district/repository"

type RepositoryPool struct {
	productRepository *repository.ProductRepository
	reviewRepository  *repository.ReviewRepository
	userRepository    *repository.UserRepository
}

func NewRepositoryPool(
	productRepository *repository.ProductRepository,
	reviewRepository *repository.ReviewRepository,
	userRepository *repository.UserRepository,
) *RepositoryPool {
	return &RepositoryPool{
		productRepository: productRepository,
		reviewRepository:  reviewRepository,
		userRepository:    userRepository,
	}
}

func (rp *RepositoryPool) GetProductRepository() *repository.ProductRepository {
	return rp.productRepository
}

func (rp *RepositoryPool) GetReviewRepository() *repository.ReviewRepository {
	return rp.reviewRepository
}

func (rp *RepositoryPool) GetUserRepository() *repository.UserRepository {
	return rp.userRepository
}
