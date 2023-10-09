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
	query := "INSERT INTO users (identification, email, username, password, address, balance) VALUES ($1, $2, $3, $4, $5, $6)"
	err := r.db.QueryRow(query, user.Email, user.Username, user.Password, user.Address, user.Balance)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT identification, email, username, password, address, balance FROM users WHERE email = $1"
	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.Identification,
		&user.Email,
		&user.Password,
		&user.Address,
		&user.Balance,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetUserByIdentification(identification int) (*models.User, error) {
	query := "SELECT identification, email, username, password, address, balance FROM users WHERE identification = $1"
	user := &models.User{}
	err := r.db.QueryRow(query, identification).Scan(
		&user.Identification,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Address,
		&user.Balance,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}
