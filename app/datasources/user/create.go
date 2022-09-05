package user

import (
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
)

// Create a new user
func (d *Datasource) Create(user entities.User) (uint64, error) {

	stmt, err := d.DB.Prepare(
		"INSERT INTO users (firstname,lastname,email,password)VALUES(?, ?, ?, ?)",
	)

	if err != nil {
		return 0, &fail.SqlInsertFailure{M: "error while preparing sql statement", E: err}
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Password,
	)

	if err != nil {
		return 0, &fail.SqlInsertFailure{M: "'error while inserting user'", E: err}
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, &fail.GetLastInsertIdFailure{M: "error while getting last insert id", E: err}
	}

	return uint64(id), nil
}
