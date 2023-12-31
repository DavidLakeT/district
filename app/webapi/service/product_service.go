package service

import (
	controller "district/controller/request"
	"district/model"
	dto "district/model/dto"
	repository "district/repository/handler"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
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

func (ps *ProductService) CreateProduct(token string, product *model.Product) (*model.Product, error) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	is_admin, err := strconv.ParseBool(tokenValues[3])
	if err != nil {
		return nil, fmt.Errorf("your session token is not valid.")
	}

	if !is_admin {
		return nil, fmt.Errorf("you have to be an administrator to create products.")
	}

	product, err = ps.repositoryPool.ProductRepository.CreateProduct(product)
	if err != nil {
		return nil, fmt.Errorf("there was an error creating the product: %v", err)
	}

	return product, nil
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

func (ps *ProductService) UpdateProduct(token string, id int, request *controller.UpdateProductRequest) error {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	is_admin, err := strconv.ParseBool(tokenValues[3])
	if err != nil {
		return fmt.Errorf("your session token is not valid.")
	}

	if !is_admin {
		return fmt.Errorf("you have to be an administrator to create products.")
	}

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

func (ps *ProductService) GetProductPicture(filename string) (*os.File, error) {
	file, err := os.Open("uploads/" + filename)
	if err != nil {
		return nil, fmt.Errorf("file was not found.")
	}
	return file, nil
}

func (ps *ProductService) UploadProductPicture(token string, file *multipart.FileHeader) error {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	is_admin, err := strconv.ParseBool(tokenValues[3])
	if err != nil {
		return fmt.Errorf("your session token is not valid.")
	}

	if !is_admin {
		return fmt.Errorf("you have to be an administrator to update product picture.")
	}

	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("file could not be read.")
	}

	dst, err := os.Create("uploads/" + file.Filename)
	if err != nil {
		return fmt.Errorf("file could not be stored.")
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProductService) DeleteProduct(token string, id int) error {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}

	tokenValues := strings.Split(string(decodedToken), ":")
	is_admin, err := strconv.ParseBool(tokenValues[3])
	if err != nil {
		return fmt.Errorf("your session token is not valid.")
	}

	if !is_admin {
		return fmt.Errorf("you have to be an administrator to delete products.")
	}

	err = ps.repositoryPool.ProductRepository.DeleteProduct(id)
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
