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

var emailRequestPasswordResetbody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Password reset</title>
</head>
<body>
   <p>Password reset link <b>Go TaskApp</b></p>
   <p>Click this link <a href="{{LINK}}">here</a> to reset your password.</p>
   <p>If you have not requested this link, ignore it.</p>
</body>
`

type sendEmailPasswordReset struct {
	Email string `form:"email" binding:"required,email"`
}

// Request password reset
func SendEmailPasswordReset(c *gin.Context) {

	var form sendEmailPasswordReset

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	credential, err := auth.SendEmailPasswordReset(form.Email)

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
		case *fail.SqlSelectNotFoundFailure:
			helpers.ApiResponse(c, false, http.StatusNotFound, "email not linked a account", nil)
			return
		default:
			helpers.ApiResponse(c, false, http.StatusInternalServerError, "an unexpected error occurred", nil)
			return
		}
	}

	link := fmt.Sprintf("http://%s%s%s", config.APP_HOST_FULL, "/auth/password/reset/", credential.Token)

	emailRequestPasswordResetbody = strings.ReplaceAll(emailRequestPasswordResetbody, "{{LINK}}", link)

	err = helpers.SendEmail([]string{credential.User.Email}, []string{}, "Password Reset", emailRequestPasswordResetbody)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponse(c, false, http.StatusInternalServerError, "Server internal error to send email verification", nil)
		return
	}

	helpers.ApiResponse(c, true, http.StatusOK, "A password reset link has been sent to your email, if it's not in your inbox check your span.", nil)
}
