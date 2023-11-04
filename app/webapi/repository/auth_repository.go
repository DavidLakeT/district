package repository

import (
	"database/sql"
	"district/model"
	"errors"
	"fmt"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) VerifyLogin(email, password string) (*model.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s' AND password = '%s'", email, password)
	fmt.Println(query)
	user := &model.User{}
	err := ar.db.QueryRow(query).Scan(
		&user.Identification,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Address,
		&user.Balance,
		&user.IsAdmin,
		&user.CreatedAt,
		&user.UpdatedAt,
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
