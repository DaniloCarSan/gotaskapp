package auth

import (
	"gotaskapp/src/database"
	"gotaskapp/src/entities"
	repositories "gotaskapp/src/repositories/user"
	"gotaskapp/src/security"
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

	db, err := database.Connect()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	repository := repositories.User(db)

	if user, err = repository.ById(id); err != nil {
		c.JSON(http.StatusUnauthorized, "Error, invalid confirmation token or it has expired, request another")
		return
	}

	if user.IsEmailVerified() {

		if err := repository.SetEmailVerified(id); err != nil {
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			c.JSON(http.StatusInternalServerError, "Server internal error")
			return
		}
	}

	c.JSON(http.StatusOK, "Email successfully confirmed, enjoy the app")
}
