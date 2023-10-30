package service

import (
	"district/model"
	repository "district/repository/handler"
	"encoding/base64"
	"encoding/json"
)

type CartService struct {
	repositoryPool *repository.RepositoryPool
}

func NewCartService(repositoryPool *repository.RepositoryPool) *CartService {
	return &CartService{repositoryPool: repositoryPool}
}

func (cs *CartService) GetCartTotal(cartToken string) (*float64, error) {
	var total float64
	cart, err := cs.decodeCartCookie(cartToken)
	if err != nil {
		return nil, err
	}

	for _, item := range cart {
		total += item.Cost * float64(item.Quantity)
	}

	return &total, nil
}

func (s *CartService) decodeCartCookie(cartToken string) (model.Cart, error) {
	var cart model.Cart

	decodedToken, err := base64.StdEncoding.DecodeString(cartToken)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(decodedToken, &cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
