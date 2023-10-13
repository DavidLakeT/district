package database

import (
	"database/sql"
)

func SetupDatabase(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS users;`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			identification INTEGER PRIMARY KEY,
			email VARCHAR(60) NOT NULL UNIQUE,
			username VARCHAR(60) NOT NULL UNIQUE,
			password VARCHAR(60) NOT NULL,
			address VARCHAR(60) NOT NULL, 
			balance FLOAT(8) DEFAULT 0.00,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW(),
			deleted_at TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
