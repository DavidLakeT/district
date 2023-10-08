package repository

import (
	"database/sql"
	"district/models"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(
		query,
		user.Username,
		user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = $1"
	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	query := "UPDATE users SET username = $1, password = $2 WHERE id = $3"
	_, err := r.db.Exec(query, user.Username, user.Password, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
