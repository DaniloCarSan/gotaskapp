package auth

import (
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/helpers"
	"gotaskapp/app/repositories/auth"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type signin struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=16"`
}

// Sign in
func SignIn(c *gin.Context) {

	var form signin

	if err := c.ShouldBind(&form); err != nil {
		helpers.ApiResponseError(c, http.StatusBadRequest, "FORM_FIELDS_INVALID", "invalid form fields", err.Error())
		return
	}

	credential, err := auth.SigIn(
		entities.Auth{
			Email:    form.Email,
			Password: form.Password,
		},
	)

	if err != nil {
		switch err.(type) {
		case *fail.DatabaseConnectFailure,
			*fail.SqlSelectFailure,
			*fail.PasswordToHashFailure,
			*fail.GenerateJwtTokenFailure:
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			helpers.ApiResponseError(c, http.StatusInternalServerError, "SERVER_INTERNAL_ERROR", "Server internal error", nil)
			return
		case *fail.SignInFailure:
			helpers.ApiResponseError(c, http.StatusUnauthorized, "EMAIL_OR_PASSWORD_INVALID", "email or password invalid", nil)
			return
		default:
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			helpers.ApiResponseError(c, http.StatusInternalServerError, "SERVER_INTERNAL_ERROR", "Server internal error", nil)
			return
		}
	}

	helpers.ApiResponseSuccess(c, http.StatusOK, credential)
}
