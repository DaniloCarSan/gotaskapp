package auth

import (
	"fmt"
	"gotaskapp/app/config"
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/helpers"
	"gotaskapp/app/repositories/auth"
	"net/http"
	"strings"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

var emailSignUpbody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Link de confirmação da conta</title>
</head>
<body>
   <p>Obrigago por criar uma conta no <b>Go TaskApp</b></p>
   <p>Clique neste link <a href="{{LINK}}">aqui</a> para confirmar seu email.</p>
   <p>Caso não tenha criado uma conta ignore este email.</p>
</body>
`

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

	credential, err := auth.SignUp(user)

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
		link := fmt.Sprintf("%s%s%s", config.APP_URL, "/auth/email/verify/", credential.Token)

		emailSignUpbody = strings.ReplaceAll(emailSignUpbody, "{{LINK}}", link)

		err = helpers.SendEmail([]string{user.Email}, []string{}, "Confirmação de email", emailSignUpbody)

		if err != nil {
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
		}
	}()

	helpers.ApiResponse(
		c,
		true,
		http.StatusOK,
		"Account created successfully, a verification link has been sent to your email.",
		nil,
	)
}
