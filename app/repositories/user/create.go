package user

import "gotaskapp/app/entities"

// Create a new user
func (r *Repository) Create(user entities.User) (uint64, error) {

	stmt, err := r.DB.Prepare(
		"INSERT INTO users (firstname,lastname,email,password)VALUES(?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Password,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}
