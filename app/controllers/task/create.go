package task

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type taskCreate struct {
	Description string `form:"description" binding:"required"`
	StatusID    uint64 `form:"status_id" binding:"required"`
}

// Create task
func Create(c *gin.Context) {

	var form taskCreate

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userID uint64
	var err error

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

	task := entities.Task{
		Description: form.Description,
		StatusID:    form.StatusID,
		UserID:      userID,
	}

	exists, err := repository.Task.ByDescriptionAndUserIDAndStatusID(task.Description, task.UserID, task.StatusID)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Task already exists"})
		return
	}

	id, err := repository.Task.Create(task)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	task.ID = id

	c.JSON(http.StatusCreated, task)
}
