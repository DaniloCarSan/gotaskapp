package user

import "gotaskapp/app/entities"

// Update user
func (r *Repository) Update(user entities.User) error {

	stmt, err := r.DB.Prepare(
		`UPDATE users SET firstname = ?, lastname = ? WHERE id = ?`,
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Firstname, user.Lastname, user.ID)

	if err != nil {
		return err
	}

	return nil
}
