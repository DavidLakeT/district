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

func (r *ReviewRepository) CreateReview(review *model.Review) error {
	query := `INSERT INTO reviews (user_id, product_id, content) VALUES ($1, $2, $3) RETURNING id`
	_, err := r.db.Exec(query, review.UserID, review.ProductID, review.Content)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create review: %v", err))
	}
	return nil
}

func (r *ReviewRepository) GetReviewById(id int) (*model.Review, error) {
	query := `SELECT id, user_id, product_id, content FROM reviews WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)
	review := &model.Review{}
	err := row.Scan(&review.ID, &review.UserID, &review.ProductID, &review.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(fmt.Sprintf("review with id %d not found", id))
		}
		return nil, errors.New(fmt.Sprintf("failed to get review: %v", err))
	}
	return review, nil
}

func (r *ReviewRepository) UpdateReview(review *model.Review) error {
	query := `UPDATE reviews SET content = $1 WHERE id = $2 AND deleted_at IS NULL`
	result, err := r.db.Exec(query, review.Content, review.ID)
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
