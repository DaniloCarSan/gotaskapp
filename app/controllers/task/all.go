package task

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"
	"strconv"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// Select all tasks
func All(c *gin.Context) {

	var userID uint64
	var err error
	var statusID uint64

	statusID, err = strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid status id")
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
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	status, err := repository.Status.ByID(statusID)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if status.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No status found"})
		return
	}

	if status.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No status found"})
		return
	}

	tasks, err := repository.Task.AllByStatusID(statusID)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusOK, []entities.Task{})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
