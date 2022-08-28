package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"
	"strings"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type passwordReset struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=16"`
}

// Password reset
func PasswordReset(c *gin.Context) {

	var form passwordReset

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString := c.Param("token")

	var token, err = security.ValidateJwtToken(tokenString)
	var id uint64

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Error, invalid confirmation token or it has expired, request another")
		return
	}

	id, err = security.ExtractUserIdOfJwtToken(token)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, err.Error())
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

	var user entities.User

	if user, err = repository.User.ById(id); err != nil {
		c.JSON(http.StatusUnauthorized, "error, invalid confirmation token or it has expired, request another")
		return
	}

	if !strings.EqualFold(user.Email, form.Email) {
		c.JSON(http.StatusUnauthorized, "the email provided does not correspond to this password reset request.")
		return
	}

	user.Password = form.Password

	err = user.PasswordToHash()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an internal error occurred, please try again later."})
		return
	}

	err = repository.User.PasswordReset(user.ID, user.Password)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "an internal error occurred, please try again later."})
		return
	}

	c.JSON(http.StatusOK, "Password successfully changed, sign in now and enjoy our features.")
}
