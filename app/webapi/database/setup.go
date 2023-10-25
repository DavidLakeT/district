package database

import (
	"database/sql"
)

func SetupDatabase(db *sql.DB) error {
	if _, err := db.Exec(`DROP TABLE IF EXISTS products CASCADE;`); err != nil {
		return err
	}
	if _, err := db.Exec(`DROP TABLE IF EXISTS reviews;`); err != nil {
		return err
	}
	if _, err := db.Exec(`DROP TABLE IF EXISTS users;`); err != nil {
		return err
	}

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			identification INTEGER PRIMARY KEY,
			email VARCHAR(60) NOT NULL UNIQUE,
			username VARCHAR(60) NOT NULL UNIQUE,
			password VARCHAR(60) NOT NULL,
			address VARCHAR(60) NOT NULL, 
			balance FLOAT(8) DEFAULT 0.00,
			is_admin BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			deleted_at TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(60) NOT NULL,
			description TEXT NOT NULL,
			stock INTEGER NOT NULL,
			price FLOAT(8) NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			deleted_at TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS reviews (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL,
		user_email VARCHAR(60) NOT NULL,
		product_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		deleted_at TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users (identification),
		FOREIGN KEY (product_id) REFERENCES products (id)
	);
`)
	if err != nil {
		return err
	}

	return nil
}
