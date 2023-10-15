package repository

import (
	"database/sql"
	"district/model"
	"errors"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2)"
	_, err := r.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func (r *ProductRepository) GetAllProducts() ([]*model.Product, error) {
	rows, err := r.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	products := make([]*model.Product, 0)
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to get products: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) GetProductByIdentification(id int) (*model.Product, error) {
	product := &model.Product{}
	err := r.db.QueryRow("SELECT * FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	query := "SELECT deleted_at FROM users WHERE identification = $1"
	var deletedAt sql.NullTime
	err := r.db.QueryRow(query, product.ID).Scan(&deletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("product not found: %w", err)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	if deletedAt.Valid {
		return fmt.Errorf("product has been deleted")
	}

	query = "UPDATE products SET name = $1, price = $2 WHERE identification = $3"
	_, err = r.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *ProductRepository) DeleteProduct(id int) error {
	_, err := r.db.Exec("UPDATE products SET deleted_at = NOW() WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
