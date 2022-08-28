package status

func (r *Repository) Create(name string, userID uint64) (uint64, error) {

	stmt, err := r.DB.Prepare("INSERT INTO status (name,user_id) VALUES(?,?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, userID)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}
