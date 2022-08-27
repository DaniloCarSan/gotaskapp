package status

import (
	"gotaskapp/app/database"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type Validate struct {
	ID uint64 `form:"id" binding:"required"`
}

// Delete status by id
func Delete(c *gin.Context) {

	var form Validate

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

	exists, err := repository.Status.ByID(form.ID)

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

	if err = repository.Status.Delete(form.ID); err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status deleted"})
}
