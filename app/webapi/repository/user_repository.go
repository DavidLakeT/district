package repository

import (
	"database/sql"
	"district/model"
	"errors"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	query := "INSERT INTO users (identification, email, username, password, address, balance) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, user.Identification, user.Email, user.Username, user.Password, user.Address, user.Balance)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	query := "SELECT identification, email, username, password, address, balance FROM users WHERE email = $1"
	user := &model.User{}
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

func (r *UserRepository) GetUserByIdentification(identification int) (*model.User, error) {
	query := "SELECT identification, email, username, password, address, balance FROM users WHERE identification = $1"
	user := &model.User{}
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

func (r *UserRepository) UpdateUser(user *model.User) error {
	query := "UPDATE users SET email = $1, username = $2, address = $3 WHERE identification = $4"
	_, err := r.db.Exec(query, user.Email, user.Username, user.Address, user.Identification)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
