package user

// Password reset
func (r *Repository) PasswordReset(id uint64, password string) error {

	stmt, err := r.DB.Prepare(
		`UPDATE users SET verified = "Y", password = ? WHERE id = ?`,
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(password, id)

	if err != nil {
		return err
	}

	return nil
}
