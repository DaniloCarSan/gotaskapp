package task

// Change status of task
func (r *Repository) ChangeStatus(statusID uint64, taskID uint64) error {

	stmt, err := r.DB.Prepare(
		"UPDATE tasks SET status_id = ? WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		statusID,
		taskID,
	)

	if err != nil {
		return err
	}

	return nil
}
