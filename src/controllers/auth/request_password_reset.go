package auth

import (
	"fmt"
	"gotaskapp/src/config"
	"gotaskapp/src/database"
	"gotaskapp/src/helpers"
	repositories "gotaskapp/src/repositories/user"
	"gotaskapp/src/security"
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
   <title>Link de redefinição de senha</title>
</head>
<body>
   <p>Lin de redefinição de senha do <b>Go TaskApp</b></p>
   <p>Clique neste link <a href="{{LINK}}">aqui</a> para redefinir a senha.</p>
   <p>Caso não tenha requisitado este link ignore.</p>
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

	db, err := database.Connect()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	repository := repositories.User(db)

	user, err := repository.ByEmail(form.Email)

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

	emailSignUpbody = strings.ReplaceAll(emailSignUpbody, "{{LINK}}", link)

	err = helpers.SendEmail([]string{user.Email}, []string{}, "Redefinição de senha", emailSignUpbody)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Um link de redefinição de senha foi enviado para seu email, aso não esteja em sua caixa de entrada verifique a de span.",
	})
}
