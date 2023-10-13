package model

import (
	"database/sql"
)

type User struct {
	Identification int     `json:"identification"`
	Email          string  `json:"email"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	Address        string  `json:"address"`
	Balance        float64 `json:"balance"`
	IsAdmin        *bool   `json:"isAdmin"`
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
	DeletedAt      sql.NullTime
}
