package auth

import (
	"gotaskapp/src/database"
	"gotaskapp/src/entities"
	"gotaskapp/src/logs"
	repositories "gotaskapp/src/repositories/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type validate struct {
	Firstname string `form:"firstname" binding:"required,alpha"`
	Lastname  string `form:"lastname" binding:"required,alpha"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required,min=6,max=16"`
}

// Sign up
func SignUp(c *gin.Context) {

	var v validate

	if err := c.ShouldBind(&v); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.Connect()

	if err != nil {
		logs.SentryCaptureAndSendException(c, err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	repository := repositories.User(db)

	user := entities.User{
		Firstname: v.Firstname,
		Lastname:  v.Lastname,
		Email:     v.Email,
		Password:  v.Password,
	}

	id, err := repository.Create(user)

	if err != nil {
		logs.SentryCaptureAndSendException(c, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Id = id

	c.JSON(http.StatusOK, user)
}
