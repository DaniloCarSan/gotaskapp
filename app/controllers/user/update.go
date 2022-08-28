package task

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type userUpdate struct {
	Firstname string `form:"firstname" binding:"required"`
	Lastname  string `form:"lastname" binding:"required"`
}

// Update ser
func Update(ctx *gin.Context) {

	var form userUpdate

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id uint64
	var err error

	if id, err = security.ExtractUserIfOFJwtTokenFromHeaderAuthorization(ctx.Request); err != nil {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.CaptureException(err)
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	repository, err := database.Repository()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.CaptureException(err)
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var user entities.User

	if user, err = repository.User.ById(id); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	user.Firstname = form.Firstname
	user.Lastname = form.Lastname

	if err = repository.User.Update(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated"})
}
