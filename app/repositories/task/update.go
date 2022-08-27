package task

// Update task
func (r *Repository) Update(id uint64, description string) error {

	stmt, err := r.DB.Prepare(
		"UPDATE tasks SET description = ? WHERE id = ? LIMIT 1",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		description,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
