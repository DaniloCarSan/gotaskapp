package status

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type statusUpdate struct {
	ID   uint64 `form:"id" binding:"required"`
	Name string `form:"name" binding:"required"`
}

// Update status by id and user id
func Update(c *gin.Context) {

	var form statusUpdate

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
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	repository, err := database.Repository()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var status entities.Status

	status.ID = form.ID
	status.Name = form.Name
	status.UserID = userID

	exists, err := repository.Status.ByID(status.ID)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}

	if exists.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}

	if err = repository.Status.Update(status); err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}
