package repository

import (
	"database/sql"
	"district/model"
	"errors"
	"fmt"
	"time"
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
	query := "SELECT identification, email, username, password, address, balance, deleted_at FROM users WHERE email = $1"
	user := &model.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.Identification,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Address,
		&user.Balance,
		&user.DeletedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user.DeletedAt.Valid {
		return nil, fmt.Errorf("user has been deleted")
	}
	return user, nil
}

func (r *UserRepository) GetUserByIdentification(identification int) (*model.User, error) {
	query := "SELECT identification, email, username, password, address, balance, deleted_at FROM users WHERE identification = $1"
	user := &model.User{}
	err := r.db.QueryRow(query, identification).Scan(
		&user.Identification,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Address,
		&user.Balance,
		&user.DeletedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user.DeletedAt.Valid {
		return nil, fmt.Errorf("user has been deleted")
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	query := "SELECT deleted_at FROM users WHERE identification = $1"
	var deletedAt sql.NullTime
	err := r.db.QueryRow(query, user.Identification).Scan(&deletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user not found: %w", err)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	if deletedAt.Valid {
		return fmt.Errorf("user has been deleted")
	}

	query = "UPDATE users SET email = $1, username = $2, address = $3 WHERE identification = $4"
	_, err = r.db.Exec(query, user.Email, user.Username, user.Address, user.Identification)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteUserByIdentification(identification int) error {
	query := "UPDATE users SET deleted_at = $1 WHERE identification = $2"
	_, err := r.db.Exec(query, time.Now().Format("2006-01-02 15:04:05"), identification)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
