package user

import "database/sql"

type Datasource struct {
	DB *sql.DB
}
