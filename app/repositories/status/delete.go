package status

// DElete status by id
func (r *Repository) Delete(id uint64) error {

	_, err := r.DB.Exec(
		"DELETE FROM status WHERE id = ? LIMIT 1",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
