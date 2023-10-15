package model

import (
	"database/sql"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
