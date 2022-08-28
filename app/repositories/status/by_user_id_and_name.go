package status

import "gotaskapp/app/entities"

// select status by user id and name
func (r *Repository) ByUserIDAndName(userID uint64, name string) (entities.Status, error) {

	rows, err := r.DB.Query(
		"SELECT * FROM status WHERE user_id = ? AND name = ? LIMIT 1",
		userID,
		name,
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
