package user

import (
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
)

// Select a user by email
func (d *Datasource) ByEmail(email string) (entities.User, error) {

	rows, err := d.DB.Query(
		"SELECT * FROM users WHERE email = ? LIMIT 1",
		email,
	)

	if err != nil {
		return entities.User{}, &fail.SqlSelectFailure{M: "error while selecting user by email", E: err}
	}

	defer rows.Close()

	if !rows.Next() {
		return entities.User{}, &fail.SqlSelectNotFoundFailure{M: "user not found", E: err}
	}

	var user entities.User

	err = rows.Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Password,
		&user.Verified,
		&user.CreateAt,
	)

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
