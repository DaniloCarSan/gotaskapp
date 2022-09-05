package auth

import (
	"fmt"
	"gotaskapp/app/config"
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/helpers"
	"gotaskapp/app/repositories/auth"
	"gotaskapp/app/security"
	"net/http"
	"strings"
	"time"

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
		helpers.ApiResponse(c, false, http.StatusBadRequest, "error in form data", gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		Firstname: form.Firstname,
		Lastname:  form.Lastname,
		Email:     form.Email,
		Password:  form.Password,
	}

	id, err := auth.SignUp(user)

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
			helpers.ApiResponse(c, false, http.StatusInternalServerError, "Server internal error", nil)
			return
		case *fail.SignUpFailure:
			helpers.ApiResponse(c, false, http.StatusBadRequest, err.Error(), nil)
			return
		}
	}

	user.ID = id

	token, err := security.GenerateJwtToken(user.ID, time.Hour*6)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponse(c, false, http.StatusInternalServerError, "Server internal error", nil)
		return
	}

	link := fmt.Sprintf("http://%s%s%s", config.APP_HOST_FULL, "/auth/email/verify/", token)

	emailSignUpbody = strings.ReplaceAll(emailSignUpbody, "{{LINK}}", link)

	err = helpers.SendEmail([]string{user.Email}, []string{}, "Confirmação de email", emailSignUpbody)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponse(c, false, http.StatusInternalServerError, "Server internal error", nil)
		return
	}

	helpers.ApiResponse(
		c,
		true,
		http.StatusOK,
		"Account created successfully, a verification link has been sent to your email.",
		nil,
	)
}
