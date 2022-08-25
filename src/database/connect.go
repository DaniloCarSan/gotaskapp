package database

import (
	"database/sql"
	"fmt"
	"gotaskapp/src/config"
	repositoryUser "gotaskapp/src/repositories/user"

	_ "github.com/go-sql-driver/mysql"
)

// Database Connect
func Repository() (*repository, error) {
	db, err := sql.Open(config.DB_DRIVE, config.DB_ADDR)

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting database: %w", err)
	}

	return &repository{
		User: &repositoryUser.Repository{DB: db},
	}, nil
}

type repository struct {
	User *repositoryUser.Repository
}
