package status

import "gotaskapp/app/entities"

// Status by id
func (r *Repository) ByID(id uint64) (entities.Status, error) {

	rows, err := r.DB.Query(
		"SELECT * FROM status WHERE id = ? LIMIT 1",
		id,
	)

	if err != nil {
		return entities.Status{}, err
	}

	defer rows.Close()

	if !rows.Next() {
		return entities.Status{}, nil
	}

	var status entities.Status

	err = rows.Scan(
		&status.ID,
		&status.Name,
		&status.UserID,
	)

	if err != nil {
		return entities.Status{}, err
	}

	return status, nil
}
