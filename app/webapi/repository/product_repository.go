package repository

import (
	"database/sql"
	"district/model"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	query := "INSERT INTO products (name, description, stock, price) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Stock, product.Price)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func (r *ProductRepository) GetAllProducts() ([]*model.Product, error) {
	rows, err := r.db.Query("SELECT * FROM products WHERE deleted_at IS NULL")
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	products := make([]*model.Product, 0)
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Stock,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		)
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

func (r *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	product := &model.Product{}
	err := r.db.QueryRow("SELECT id, name, description, stock, price FROM products WHERE id = $1 AND deleted_at IS NULL", id).Scan(&product.ID, &product.Name, &product.Description, &product.Stock, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
}

func (r *ProductRepository) GetProductsByName(name string) ([]*model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, stock, price FROM products WHERE name LIKE '%' || $1 || '%' AND deleted_at IS NULL", name)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}
	defer rows.Close()

	products := make([]*model.Product, 0)
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Stock, &product.Price)
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

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	query := "UPDATE products SET name = $1, price = $2 WHERE id = $3 AND deleted_at IS NULL"
	result, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found or has been deleted")
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
