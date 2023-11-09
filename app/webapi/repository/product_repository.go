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

func (pr *ProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	query := "INSERT INTO products (name, description, stock, price) VALUES ($1, $2, $3, $4) RETURNING id"
	err := pr.db.QueryRow(query, product.Name, product.Description, product.Stock, product.Price).Scan(&product.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

func (pr *ProductRepository) GetAllProducts() ([]*model.Product, error) {
	rows, err := pr.db.Query("SELECT * FROM products WHERE deleted_at IS NULL")
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

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	product := &model.Product{}
	err := pr.db.QueryRow("SELECT id, name, description, stock, price FROM products WHERE id = $1 AND deleted_at IS NULL", id).Scan(&product.ID, &product.Name, &product.Description, &product.Stock, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
}

func (pr *ProductRepository) GetProductsByName(name string) ([]*model.Product, error) {
	rows, err := pr.db.Query("SELECT id, name, description, stock, price FROM products WHERE name LIKE '%' || $1 || '%' AND deleted_at IS NULL", name)
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

func (pr *ProductRepository) UpdateProduct(product *model.Product) error {
	query := "UPDATE products SET name = $1, description = $2, stock = $3, price = $4 WHERE id = $5 AND deleted_at IS NULL"
	result, err := pr.db.Exec(query, product.Name, product.Description, product.Stock, product.Price, product.ID)
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

func (pr *ProductRepository) DeleteProduct(id int) error {
	result, err := pr.db.Exec("UPDATE products SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL", id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found or has already been deleted")
	}

	return nil
}

func (pr *ProductRepository) CreateProductsTable() error {
	_, err := pr.db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(60) NOT NULL,
			description TEXT NOT NULL,
			stock INTEGER NOT NULL,
			price FLOAT(8) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			deleted_at TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepository) DeleteProductsTable() error {
	_, err := pr.db.Exec("DROP TABLE IF EXISTS products CASCADE")
	if err != nil {
		return fmt.Errorf("failed to delete products table: %w", err)
	}
	return nil
}
