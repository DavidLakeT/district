package repository

import (
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

/*

func (r *AuthRepository) Signup(user *models.User) (*models.User, error) {
	_, err := r.db.Exec("INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)", user.Username, string(hashedPassword), time.Now())
	if err != nil {
		return nil, err
	}

	return nil
}

func (r *AuthRepository) Login(username, password string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}

*/
