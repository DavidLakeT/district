package service

import (
	"district/model"
	dto "district/model/dto"
	"district/repository"
)

type ProductService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (p *ProductService) GetAllProducts() ([]*dto.ProductDTO, error) {
	products, err := p.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	productDTOs := make([]*dto.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = dto.ConvertToProductDTO(product)
	}

	return productDTOs, nil
}

func (p *ProductService) CreateProduct(product *model.Product) error {
	return p.productRepository.CreateProduct(product)
}

func (p *ProductService) GetProductById(id int) (*dto.ProductDTO, error) {
	product, err := p.productRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	productDTO := dto.ConvertToProductDTO(product)

	return productDTO, nil
}

func (p *ProductService) GetProductsByName(name string) ([]*dto.ProductDTO, error) {
	products, err := p.productRepository.GetProductsByName(name)
	if err != nil {
		return nil, err
	}

	productDTOs := make([]*dto.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = dto.ConvertToProductDTO(product)
	}

	return productDTOs, nil
}
