package task

// Delete task by ID
func (r *Repository) Delete(taskID uint64) error {

	_, err := r.DB.Exec(
		"DELETE FROM tasks WHERE id = ? LIMIT 1",
		taskID,
	)

	if err != nil {
		return err
	}

	return nil
}
