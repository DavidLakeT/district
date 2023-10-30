package service

import repository "district/repository/handler"

type CartService struct {
	repositoryPool *repository.RepositoryPool
}

func NewCartService(repositoryPool *repository.RepositoryPool) *CartService {
	return &CartService{repositoryPool: repositoryPool}
}
