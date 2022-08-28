package database

import (
	"database/sql"
	"fmt"
	"gotaskapp/app/config"
	repositoryStatus "gotaskapp/app/repositories/status"
	repositoryTask "gotaskapp/app/repositories/task"
	repositoryUser "gotaskapp/app/repositories/user"

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
		User:   &repositoryUser.Repository{DB: db},
		Status: &repositoryStatus.Repository{DB: db},
		Task:   &repositoryTask.Repository{DB: db},
	}, nil
}

type repository struct {
	User   *repositoryUser.Repository
	Status *repositoryStatus.Repository
	Task   *repositoryTask.Repository
}
