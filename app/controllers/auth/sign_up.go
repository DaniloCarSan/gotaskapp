package auth

import (
	"bytes"
	"encoding/json"
	"gotaskapp/app/config"
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/helpers"
	"gotaskapp/app/repositories/auth"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type signUp struct {
	Firstname string `form:"firstname" binding:"required,alpha"`
	Lastname  string `form:"lastname" binding:"required,alpha"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required,min=6,max=16"`
}

// Sign up
func SignUp(c *gin.Context) {

	var form signUp

	if err := c.ShouldBind(&form); err != nil {
		helpers.ApiResponseError(c, http.StatusBadRequest, "FORM_VALIDATE_ERROR", err.Error(), nil)
		return
	}

	user := entities.User{
		Firstname: form.Firstname,
		Lastname:  form.Lastname,
		Email:     form.Email,
		Password:  form.Password,
	}

	_, err := auth.SignUp(user)

	if err != nil {
		switch err.(type) {
		case *fail.DatabaseConnectFailure,
			*fail.SqlInsertFailure,
			*fail.GenerateJwtTokenFailure,
			*fail.PasswordToHashFailure,
			*fail.GetLastInsertIdFailure:
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			helpers.ApiResponseError(c, http.StatusInternalServerError, "SERVER_INTERNAL_ERROR", "Internal server error", nil)
			return
		case *fail.SignUpFailure:
			helpers.ApiResponseError(c, http.StatusBadRequest, "SIGN_UP_ERROR", err.Error(), nil)
			return
		default:
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			helpers.ApiResponseError(c, http.StatusInternalServerError, "SERVER_INTERNAL_ERROR", "an unexpected error occurred", nil)
			return
		}
	}

	go func() {
		// Send email verification
		body, _ := json.Marshal(
			map[string]string{
				"email": form.Email,
			},
		)
		payload := bytes.NewBuffer(body)

		http.Post(config.APP_URL+"/auth/send/email/verification", "application/json", payload)
	}()

	helpers.ApiResponse(
		c,
		true,
		http.StatusOK,
		"Account created successfully, a verification link has been sent to your email.",
		nil,
	)
}
