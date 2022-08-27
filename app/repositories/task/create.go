package task

import "gotaskapp/app/entities"

// Create task
func (r *Repository) Create(task entities.Task) (uint64, error) {

	stmt, err := r.DB.Prepare(
		"INSERT INTO tasks (description, status_id, user_id) VALUES (?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		task.Description,
		task.StatusID,
		task.UserID,
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
