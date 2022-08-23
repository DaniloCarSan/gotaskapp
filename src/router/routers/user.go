package routers

import (
	userController "gotaskapp/src/controllers/user"
	"gotaskapp/src/middlewares"

	"github.com/gin-gonic/gin"
)

// Routers of user
func User(app *gin.Engine) {
	usersGroup := app.Group("/users", middlewares.Authenticate())
	{
		usersGroup.GET("", userController.Select)
		usersGroup.PUT("", userController.Update)
	}
}
