package auth

import (
	"gotaskapp/src/database"
	repositories "gotaskapp/src/repositories/user"
	"gotaskapp/src/security"
	"net/http"
	"time"

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

	err = security.CompareHashWithPassword(user.Password, form.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password inválid"})
		return
	}

	if !user.IsEmailVerified() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email no verified"})
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

	c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}
