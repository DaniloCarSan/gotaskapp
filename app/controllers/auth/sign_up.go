package auth

import (
	"fmt"
	"gotaskapp/app/config"
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/helpers"
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

	user := entities.User{
		Firstname: form.Firstname,
		Lastname:  form.Lastname,
		Email:     form.Email,
		Password:  form.Password,
	}

	exists, err := repository.User.ByEmail(user.Email)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if exists != (entities.User{}) {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "there is already an account linked to this email"})
		return
	}

	err = user.PasswordToHash()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	id, err := repository.User.Create(user)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Id = id

	token, err := security.GenerateJwtToken(user.Id, time.Hour*6)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	link := fmt.Sprintf("http://%s%s%s", config.APP_HOST_FULL, "/auth/email/verify/", token)

	emailSignUpbody = strings.ReplaceAll(emailSignUpbody, "{{LINK}}", link)

	err = helpers.SendEmail([]string{user.Email}, []string{}, "Confirmação de email", emailSignUpbody)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account created successfully, a verification link has been sent to your email.",
	})
}
