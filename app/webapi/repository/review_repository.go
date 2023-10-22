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
