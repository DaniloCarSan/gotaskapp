package task

import "gotaskapp/app/entities"

// By id
func (r *Repository) ByID(id uint64) (entities.Task, error) {

	rows, err := r.DB.Query(
		"SELECT * FROM tasks WHERE id = ? LIMIT 1",
		id,
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
