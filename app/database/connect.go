package database

import (
	"database/sql"
	"fmt"
	"gotaskapp/app/config"
	datasourceUser "gotaskapp/app/datasources/user"
	fail "gotaskapp/app/failures"
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

// Database Connect
func Datasources() (*datasource, error) {

	db, err := sql.Open(config.DB_DRIVE, config.DB_ADDR)

	if err != nil {
		return nil, &fail.DatabaseConnectFailure{M: "error connecting database", E: err}
	}

	if err = db.Ping(); err != nil {
		return nil, &fail.DatabaseConnectFailure{M: "error connecting database", E: err}
	}

	return &datasource{
		User: &datasourceUser.Datasource{DB: db},
	}, nil
}

type datasource struct {
	User *datasourceUser.Datasource
}
