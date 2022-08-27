package status

import "gotaskapp/app/entities"

// Select all status from user id
func (r *Repository) ByUserID(userID uint64) ([]entities.Status, error) {

	var status []entities.Status

	rows, err := r.DB.Query("SELECT * FROM status WHERE user_id = ?", userID)

	if err != nil {
		return status, err
	}

	defer rows.Close()

	if !rows.Next() {
		return status, nil
	}

	for rows.Next() {

		var s entities.Status

		if err := rows.Scan(&s.ID, &s.Name, &s.UserID); err != nil {
			return status, err
		}

		status = append(status, s)
	}

	return status, nil
}
