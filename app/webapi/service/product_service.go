package service

import (
	"district/model"
	dto "district/model/dto"
	repository "district/repository/handler"
)

type ProductService struct {
	repositoryPool *repository.RepositoryPool
}

func NewProductService(repositoryPool *repository.RepositoryPool) *ProductService {
	return &ProductService{repositoryPool: repositoryPool}
}

func (ps *ProductService) GetAllProducts() ([]*dto.ProductDTO, error) {
	products, err := ps.repositoryPool.ProductRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	productDTOs := make([]*dto.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = dto.ConvertToProductDTO(product)
	}

	return productDTOs, nil
}

func (ps *ProductService) CreateProduct(product *model.Product) error {
	return ps.repositoryPool.ProductRepository.CreateProduct(product)
}

func (ps *ProductService) GetProductById(id int) (*dto.ProductDTO, error) {
	product, err := ps.repositoryPool.ProductRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	productDTO := dto.ConvertToProductDTO(product)

	return productDTO, nil
}

func (ps *ProductService) GetProductsByName(name string) ([]*dto.ProductDTO, error) {
	products, err := ps.repositoryPool.ProductRepository.GetProductsByName(name)
	if err != nil {
		return nil, err
	}

	productDTOs := make([]*dto.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = dto.ConvertToProductDTO(product)
	}

	return productDTOs, nil
}
