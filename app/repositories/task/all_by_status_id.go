package task

import "gotaskapp/app/entities"

// Select all tasks by status id
func (r *Repository) AllByStatusID(statusID uint64) ([]entities.Task, error) {

	rows, err := r.DB.Query(
		"SELECT * FROM tasks WHERE status_id = ?",
		statusID,
	)

	if err != nil {
		return []entities.Task{}, err
	}

	defer rows.Close()

	var tasks []entities.Task

	for rows.Next() {

		var task entities.Task

		err = rows.Scan(
			&task.ID,
			&task.Description,
			&task.StatusID,
			&task.UserID,
		)

		if err != nil {
			return []entities.Task{}, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
