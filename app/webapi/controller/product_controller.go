package controller

import (
	controller "district/controller/request"
	models "district/model"
	"district/service"
	"fmt"
	"net/http"

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
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}

// Endpoint: POST /api/product
// - Creates a new product with the specified information.
func (pc *ProductController) CreateProduct(c echo.Context) error {
	var request controller.CreateProductRequest

	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "Invalid product information")
	}

	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	if err := pc.productService.CreateProduct(&product); err != nil {
		fmt.Println("error:", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
	})
}

// Endpoint: GET /api/product/:name
// - Retrieves information about the specified product (name, description, price, etc).
func (pc *ProductController) SearchProducts(c echo.Context) error {
	name := c.Param("name")

	products, err := pc.productService.GetProductsByName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}
