package status

import (
	"gotaskapp/app/database"
	"gotaskapp/app/security"
	"net/http"
	"strconv"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// Delete status by id
func Delete(c *gin.Context) {

	var statusID uint64
	var userID uint64
	var err error

	statusID, err = strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
		return
	}

	userID, err = security.ExtractUserIfOFJwtTokenFromHeaderAuthorization(c.Request)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	repository, err := database.Repository()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	exists, err := repository.Status.ByID(statusID)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if exists.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}

	if err = repository.Status.Delete(statusID); err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status deleted"})
}
