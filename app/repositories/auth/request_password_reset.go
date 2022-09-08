package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/security"
	"time"
)

func RequestPasswordReset(email string) (entities.Credential, error) {

	datasources, err := database.Datasources()

	if err != nil {
		return entities.Credential{}, err
	}

	user, err := datasources.User.ByEmail(email)

	if err != nil {
		return entities.Credential{}, err
	}

	token, err := security.GenerateJwtToken(user.ID, time.Hour*6)

	if err != nil {
		return entities.Credential{}, &fail.GenerateJwtTokenFailure{M: "error generate jwt token", E: err}
	}

	return entities.Credential{
		User:  user,
		Token: token,
	}, nil
}
