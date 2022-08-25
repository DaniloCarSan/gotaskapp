package user

import "gotaskapp/src/entities"

// ByID
func (r *Repository) ById(id uint64) (entities.User, error) {

	rows, err := r.DB.Query(
		"SELECT * FROM users WHERE id = ? LIMIT 1",
		id,
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
