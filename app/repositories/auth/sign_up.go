package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/security"
	"time"
)

func SignUp(user entities.User) (entities.Credential, error) {

	datasources, err := database.Datasources()

	if err != nil {
		return entities.Credential{}, err
	}

	u, err := datasources.User.ByEmail(user.Email)

	if err != nil {
		if _, ok := err.(*fail.SqlSelectNotFoundFailure); !ok {
			return entities.Credential{}, err
		}
	}

	if u.ID > 0 {
		return entities.Credential{}, &fail.SignUpFailure{M: "there is already an account linked to this email", E: err}
	}

	err = user.PasswordToHash()

	if err != nil {
		return entities.Credential{}, &fail.PasswordToHashFailure{M: "error while hashing password", E: err}
	}

	id, err := datasources.User.Create(user)

	if err != nil {
		return entities.Credential{}, err
	}

	user.ID = id

	token, err := security.GenerateJwtToken(user.ID, time.Hour*6)

	if err != nil {
		return entities.Credential{}, &fail.GenerateJwtTokenFailure{M: "error generate jwt token", E: err}
	}

	return entities.Credential{
		User:  user,
		Token: token,
	}, nil
}
