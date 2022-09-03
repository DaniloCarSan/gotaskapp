package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/security"
	"time"
)

func SigIn(auth entities.Auth) (entities.Credential, error) {

	datasources, err := database.Datasources()

	if err != nil {
		return entities.Credential{}, err
	}

	user, err := datasources.User.ByEmail(auth.Email)

	if err != nil {
		return entities.Credential{}, err
	}

	err = security.CompareHashWithPassword(user.Password, auth.Password)

	if err != nil {
		return entities.Credential{}, &fail.SignInFailure{M: "password invalid", E: err}
	}

	if !user.IsEmailVerified() {
		return entities.Credential{}, &fail.SignInFailure{M: "email no verified", E: err}
	}

	token, err := security.GenerateJwtToken(user.ID, time.Hour*365)

	if err != nil {
		return entities.Credential{}, &fail.GenerateJwtTokenFailure{M: "error generate jwt token", E: err}
	}

	return entities.Credential{
		User:  user,
		Token: token,
	}, nil
}
