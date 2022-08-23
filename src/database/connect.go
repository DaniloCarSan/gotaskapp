package database

import (
	"database/sql"
	"gotaskapp/src/config"

	_ "github.com/go-sql-driver/mysql"
)

// Database Connect
func Connect() (*sql.DB, error) {
	db, err := sql.Open(config.DB_DRIVE, config.DB_ADDR)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
