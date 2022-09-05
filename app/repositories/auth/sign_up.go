package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
)

func SignUp(user entities.User) (uint64, error) {

	datasources, err := database.Datasources()

	if err != nil {
		return 0, err
	}

	u, err := datasources.User.ByEmail(user.Email)

	if err != nil {
		if _, ok := err.(*fail.SqlSelectNotFoundFailure); !ok {
			return 0, err
		}
	}

	if u.ID > 0 {
		return 0, &fail.SignUpFailure{M: "there is already an account linked to this email", E: err}
	}

	err = user.PasswordToHash()

	if err != nil {
		return 0, &fail.PasswordToHashFailure{M: "error while hashing password", E: err}
	}

	id, err := datasources.User.Create(user)

	if err != nil {
		return 0, err
	}

	return id, nil
}
