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
	query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price)
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
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get products: %w", err)
		}
		reviews, err := r.getProductReviews(product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get reviews for product %d: %w", product.ID, err)
		}
		product.Reviews = reviews
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) GetProductsByName(name string) ([]*model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price FROM products WHERE name LIKE '%' || $1 || '%' AND deleted_at IS NULL", name)
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
		reviews, err := r.getProductReviews(product.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get reviews for product %d: %w", product.ID, err)
		}
		product.Reviews = reviews
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return products, nil
}

func (r *ProductRepository) getProductReviews(productID int) ([]*model.Review, error) {
	rows, err := r.db.Query("SELECT id, product_id, user_id, content FROM reviews WHERE product_id = $1", productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get reviews: %w", err)
	}
	defer rows.Close()

	reviews := make([]*model.Review, 0)
	for rows.Next() {
		review := &model.Review{}
		err := rows.Scan(&review.ID, &review.ProductID, &review.AuthorID, &review.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to get reviews: %w", err)
		}
		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get reviews: %w", err)
	}

	return reviews, nil
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
