package service

import (
	controller "district/controller/request"
	"district/model"
	dto "district/model/dto"
	repository "district/repository/handler"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
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
		reviews, err := ps.repositoryPool.ReviewRepository.GetProductReviews(product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get reviews for product %d: %w", product.ID, err)
		}

		product.Reviews = reviews
		productDTOs[i] = dto.ConvertToProductDTO(product)
	}

	return productDTOs, nil
}

func (ps *ProductService) CreateProduct(token string, product *model.Product) error {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	is_admin, err := strconv.ParseBool(tokenValues[3])
	if err != nil {
		return nil
	}

	if !is_admin {
		return fmt.Errorf("you have to be an administrator to create products")
	}

	return ps.repositoryPool.ProductRepository.CreateProduct(product)
}

func (ps *ProductService) GetProductById(id int) (*dto.ProductDTO, error) {
	product, err := ps.repositoryPool.ProductRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	reviews, err := ps.repositoryPool.ReviewRepository.GetProductReviews(product.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get reviews for product %d: %w", product.ID, err)
	}
	product.Reviews = reviews

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
		reviews, err := ps.repositoryPool.ReviewRepository.GetProductReviews(product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get reviews for product %d: %w", product.ID, err)
		}
		product.Reviews = reviews

		productDTOs[i] = dto.ConvertToProductDTO(product)
	}

	return productDTOs, nil
}

func (ps *ProductService) UpdateProduct(id int, request *controller.UpdateProductRequest) error {
	product, err := ps.repositoryPool.ProductRepository.GetProductByID(id)
	if err != nil {
		return err
	}

	if request.Name != nil {
		product.Name = *request.Name
	}
	if request.Description != nil {
		product.Description = *request.Description
	}
	if request.Stock != nil {
		product.Stock = *request.Stock
	}
	if request.Price != nil {
		product.Price = *request.Price
	}

	return ps.repositoryPool.ProductRepository.UpdateProduct(product)
}

func (ps *ProductService) DeleteProduct(id int) error {
	err := ps.repositoryPool.ProductRepository.DeleteProduct(id)
	if err != nil {
		return err
	}

	err = ps.repositoryPool.ReviewRepository.DeleteReviewsByProductID(id)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) DeleteProductsTable() error {
	if err := ps.repositoryPool.ProductRepository.DeleteProductsTable(); err != nil {
		return err
	}

	return nil
}
