package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// Email verify
func EmailVerify(c *gin.Context) {

	tokenString := c.Param("token")

	var token, err = security.ValidateJwtToken(tokenString)
	var id uint64
	var user entities.User

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

	if user, err = repository.User.ById(id); err != nil {
		c.JSON(http.StatusUnauthorized, "Error, invalid confirmation token or it has expired, request another")
		return
	}

	if user.IsEmailVerified() {

		if err := repository.User.SetEmailVerified(id); err != nil {
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			c.JSON(http.StatusInternalServerError, "Server internal error")
			return
		}
	}

	c.JSON(http.StatusOK, "Email successfully confirmed, enjoy the app")
}
