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

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT username, password FROM users WHERE username = $1"
	user := &models.User{}
	err := r.db.QueryRow(query, username).Scan(&user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// TO-DO: Fix logic on username handling
func (r *UserRepository) UpdateUser(user *models.User) error {
	query := "UPDATE users SET username = $1, password = $2 WHERE username = $3"
	_, err := r.db.Exec(query, user.Username, user.Password, user.Username)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(username string) error {
	query := "DELETE FROM users WHERE username = $1"
	_, err := r.db.Exec(query, username)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
