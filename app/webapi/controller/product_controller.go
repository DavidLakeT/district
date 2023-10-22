package controller

import (
	request "district/controller/request"
	models "district/model"
	"district/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// Endpoint: GET /api/products
// - Retrieves information about all products (name, description, price, etc).
func (pc *ProductController) GetAllProducts(c echo.Context) error {
	products, err := pc.productService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, products)
}

// Endpoint: POST /api/product
// - Creates a new product with the specified information.
func (pc *ProductController) CreateProduct(c echo.Context) error {
	var request request.CreateProductRequest

	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid product information",
		})
	}

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	if err := pc.productService.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
	})
}

// Endpoint: GET /api/product/id/:id
// - Retrieves information about the specified product (name, description, price, etc).
func (pc *ProductController) SearchProductsById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid product ID",
		})
	}

	product, err := pc.productService.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}

// Endpoint: GET /api/product/name/:name
// - Retrieves information about the specified product (name, description, price, etc).
func (pc *ProductController) SearchProductsByName(c echo.Context) error {
	name := c.Param("name")

	products, err := pc.productService.GetProductsByName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, products)
}
