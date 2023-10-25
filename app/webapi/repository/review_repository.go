package repository

import (
	"database/sql"
	"district/model"
	"errors"
	"fmt"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (rr *ReviewRepository) CreateReview(review *model.Review) error {
	query := `INSERT INTO reviews (user_id, user_email, product_id, content) VALUES ($1, $2, $3, $4)`
	_, err := rr.db.Exec(query, review.UserID, review.UserEmail, review.ProductID, review.Content)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create review: %v", err))
	}
	return nil
}

func (rr *ReviewRepository) GetReviewById(id int) (*model.Review, error) {
	query := `SELECT id, user_id, user_email, product_id, content FROM reviews WHERE id = $1 AND deleted_at IS NULL`
	row := rr.db.QueryRow(query, id)
	review := &model.Review{}
	err := row.Scan(&review.ID, &review.UserID, &review.UserEmail, &review.ProductID, &review.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("review with id %d not found", id))
		}
		return nil, errors.New(fmt.Sprintf("failed to get review: %v", err))
	}
	return review, nil
}

func (rr *ReviewRepository) UpdateReview(review *model.Review) error {
	query := `UPDATE reviews SET content = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := rr.db.Exec(query, review.Content, review.ID)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update review: %v", err))
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get rows affected: %v", err))
	}
	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("review with id %d not found", review.ID))
	}
	return nil
}

func (rr *ReviewRepository) DeleteReview(id int) error {
	query := `UPDATE reviews SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := rr.db.Exec(query, id)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to delete review: %v", err))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get rows affected: %v", err))
	}
	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("review with id %d not found", id))
	}

	return nil
}

func (rr *ReviewRepository) GetProductReviews(productID int) ([]*model.Review, error) {
	rows, err := rr.db.Query("SELECT id, product_id, user_id, user_email, content FROM reviews WHERE product_id = $1", productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get reviews: %w", err)
	}
	defer rows.Close()

	reviews := make([]*model.Review, 0)
	for rows.Next() {
		review := &model.Review{}
		err := rows.Scan(&review.ID, &review.ProductID, &review.UserID, &review.UserEmail, &review.Content)
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
