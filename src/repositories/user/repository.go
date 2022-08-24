package user

import "database/sql"

type user struct {
	db *sql.DB
}

// Fabric a new repository of the user
func User(db *sql.DB) *user {
	return &user{
		db: db,
	}
}
