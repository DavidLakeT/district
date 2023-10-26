package controller

import (
	request "district/controller/request"
	models "district/model"
	service "district/service/handler"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	servicePool *service.ServicePool
}

func NewProductController(servicePool *service.ServicePool) *ProductController {
	return &ProductController{servicePool: servicePool}
}

// Endpoint: GET /api/products
// - Retrieves information about all products (name, description, price, etc).
func (pc *ProductController) GetAllProducts(c echo.Context) error {
	products, err := pc.servicePool.ProductService.GetAllProducts()
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
		Stock:       request.Stock,
		Price:       request.Price,
	}

	if err := pc.servicePool.ProductService.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"stock":       product.Stock,
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

	product, err := pc.servicePool.ProductService.GetProductById(id)
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

	products, err := pc.servicePool.ProductService.GetProductsByName(name)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, products)
}

// Endpoint: PUT /api/product/:id
// - Updates the product with the specified ID.
func (pc *ProductController) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid product ID",
		})
	}

	var request request.UpdateProductRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid request payload",
		})
	}

	err = pc.servicePool.ProductService.UpdateProduct(id, &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Product successfully updated."})
}

// Endpoint: DELETE /api/product/:id
// - Deletes the product with the specified ID (and its reviews).
func (pc *ProductController) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid product ID",
		})
	}

	err = pc.servicePool.ProductService.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "product succesfully deleted."})
}
