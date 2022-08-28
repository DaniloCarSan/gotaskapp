package status

import "gotaskapp/app/entities"

// Update status by id
func (r *Repository) Update(status entities.Status) error {

	stmt, err := r.DB.Prepare("UPDATE status SET name = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(status.Name, status.ID)

	if err != nil {
		return err
	}

	return nil
}
