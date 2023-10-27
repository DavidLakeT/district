package service

import (
	repository "district/repository/handler"
)

type UtilsService struct {
	repositoryPool *repository.RepositoryPool
}

func NewUtilsService(repositoryPool *repository.RepositoryPool) *UtilsService {
	return &UtilsService{repositoryPool: repositoryPool}
}

func (us *UtilsService) ClearDatabase() error {
	if err := us.repositoryPool.UserRepository.DeleteUsersTable(); err != nil {
		return err
	}
	if err := us.repositoryPool.UserRepository.CreateUsersTable(); err != nil {
		return err
	}

	if err := us.repositoryPool.ProductRepository.DeleteProductsTable(); err != nil {
		return err
	}
	if err := us.repositoryPool.ProductRepository.CreateProductsTable(); err != nil {
		return err
	}

	if err := us.repositoryPool.ReviewRepository.DeleteReviewsTable(); err != nil {
		return err
	}
	if err := us.repositoryPool.ReviewRepository.CreateReviewsTable(); err != nil {
		return err
	}

	return nil
}
