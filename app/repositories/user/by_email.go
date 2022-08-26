package user

import "gotaskapp/app/entities"

// Select a user by email
func (r *Repository) ByEmail(email string) (entities.User, error) {

	rows, err := r.DB.Query(
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
