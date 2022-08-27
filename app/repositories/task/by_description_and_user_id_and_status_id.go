package task

import "gotaskapp/app/entities"

// By description and user id and status id
func (r *Repository) ByDescriptionAndUserIDAndStatusID(description string, userID uint64, statusID uint64) (entities.Task, error) {

	rows, err := r.DB.Query(
		"SELECT * FROM tasks WHERE description = ? AND user_id = ? AND status_id = ? LIMIT 1",
		description,
		userID,
		statusID,
	)

	if err != nil {
		return entities.Task{}, err
	}

	defer rows.Close()

	if !rows.Next() {
		return entities.Task{}, nil
	}

	var task entities.Task

	err = rows.Scan(
		&task.ID,
		&task.Description,
		&task.StatusID,
		&task.UserID,
	)

	if err != nil {
		return entities.Task{}, err
	}

	return task, nil
}
