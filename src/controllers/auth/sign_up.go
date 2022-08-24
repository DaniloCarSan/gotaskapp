package auth

import (
	"gotaskapp/src/database"
	"gotaskapp/src/entities"
	repositories "gotaskapp/src/repositories/user"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type signUp struct {
	Firstname string `form:"firstname" binding:"required,alpha"`
	Lastname  string `form:"lastname" binding:"required,alpha"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required,min=6,max=16"`
}

// Sign up
func SignUp(c *gin.Context) {

	var form signUp

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	user := entities.User{
		Firstname: form.Firstname,
		Lastname:  form.Lastname,
		Email:     form.Email,
		Password:  form.Password,
	}

	exists, err := repository.ByEmail(user.Email)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if exists != (entities.User{}) {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "there is already an account linked to this email"})
		return
	}

	err = user.PasswordToHash()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	id, err := repository.Create(user)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Id = id

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
