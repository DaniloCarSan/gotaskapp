package user

//  Set email verified
func (r *Repository) SetEmailVerified(id uint64) error {

	stmt, err := r.DB.Prepare(
		`UPDATE users SET verified = "Y" WHERE id = ?`,
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
