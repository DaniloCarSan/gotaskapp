package auth

import (
	"fmt"
	"gotaskapp/app/config"
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
   <p>Clique neste link <a href="{{LINK}}">aqui</a> para confirmar seu email.</p>
   <p>Caso não tenha criado uma conta ignore este email.</p>
</body>
`

type sendEmailVerification struct {
	Email string `form:"email" binding:"required,email"`
}

func SendEmailVerification(c *gin.Context) {

	var form sendEmailVerification

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credential, err := auth.SendEmailVerification(form.Email)

	if err != nil {
		switch err.(type) {
		case *fail.DatabaseConnectFailure,
			*fail.SqlSelectFailure,
			*fail.GenerateJwtTokenFailure:
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			helpers.ApiResponse(c, false, http.StatusInternalServerError, "Server internal error", nil)
			return
		case *fail.SendEmailVerificationFailure:
			helpers.ApiResponse(c, false, http.StatusBadRequest, err.Error(), nil)
			return
		case *fail.SqlSelectNotFoundFailure:
			helpers.ApiResponse(c, false, http.StatusNotFound, "email not linked a account", nil)
			return
		default:
			helpers.ApiResponse(c, false, http.StatusInternalServerError, "an unexpected error occurred", nil)
			return
		}
	}

	go func() {
		link := fmt.Sprintf("%s%s%s", config.APP_URL, "/auth/email/verify/", credential.Token)

		emailSignUpbody = strings.ReplaceAll(emailSignUpbody, "{{LINK}}", link)

		err = helpers.SendEmail([]string{credential.User.Email}, []string{}, "Confirmação de email", emailSignUpbody)

		if err != nil {
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
		}
	}()

	helpers.ApiResponse(c, true, http.StatusOK, "A link has been sent to your email, if it's not in your inbox check your box span.", nil)
}
