package status

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// All status from user authenticated passed in token
func All(c *gin.Context) {

	var userID uint64
	var err error

	userID, err = security.ExtractUserIfOFJwtTokenFromHeaderAuthorization(c.Request)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	repository, err := database.Repository()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var status []entities.Status

	if status, err = repository.Status.ByUserID(userID); err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, status)
}
