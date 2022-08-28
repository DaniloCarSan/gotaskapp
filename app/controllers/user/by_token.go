package task

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// Select user by token jwt
func ByToken(c *gin.Context) {

	var id uint64
	var err error

	if id, err = security.ExtractUserIfOFJwtTokenFromHeaderAuthorization(c.Request); err != nil {
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

	c.JSON(http.StatusOK, user)
}
