package user

import "gotaskapp/src/entities"

// Select a user by email
func (repository *user) ByEmail(email string) (entities.User, error) {

	rows, err := repository.db.Query(
		"SELECT * FROM users WHERE email = ? LIMIT 1",
		email,
	)

	if err != nil {
		return entities.User{}, err
	}

	defer rows.Close()

	if !rows.Next() {
		return entities.User{}, nil
	}

	var user entities.User

	err = rows.Scan(
		&user.Id,
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
