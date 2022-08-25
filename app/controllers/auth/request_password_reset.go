package auth

import (
	"fmt"
	"gotaskapp/app/config"
	"gotaskapp/app/database"
	"gotaskapp/app/helpers"
	"gotaskapp/app/security"
	"net/http"
	"strings"
	"time"

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

type requestPasswordReset struct {
	Email string `form:"email" binding:"required,email"`
}

// Request password reset
func RequestPasswordReset(c *gin.Context) {

	var form requestPasswordReset

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repository, err := database.Repository()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	user, err := repository.User.ByEmail(form.Email)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	token, err := security.GenerateJwtToken(user.Id, time.Hour*6)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	link := fmt.Sprintf("http://%s%s%s", config.APP_HOST_FULL, "/password/reset/", token)

	emailRequestPasswordResetbody = strings.ReplaceAll(emailRequestPasswordResetbody, "{{LINK}}", link)

	err = helpers.SendEmail([]string{user.Email}, []string{}, "Password Reset", emailRequestPasswordResetbody)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "A password reset link has been sent to your email, if it's not in your inbox check your span.",
	})
}
