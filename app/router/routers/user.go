package routers

import (
	userController "gotaskapp/app/controllers/user"
	"gotaskapp/app/middlewares"

	"github.com/gin-gonic/gin"
)

// Routers of user
func User(app *gin.Engine) {
	usersGroup := app.Group("/users", middlewares.Authenticate())
	{
		usersGroup.GET("", userController.ByToken)
		usersGroup.PUT("", userController.Update)
	}
}
