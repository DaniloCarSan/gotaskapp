package router

import (
	"gotaskapp/app/router/routers"

	"github.com/gin-gonic/gin"
)

// Load all routes of the application
func LoadRouters(app *gin.Engine) *gin.Engine {

	routers.Auth(app)
	routers.User(app)
	routers.Task(app)

	return app
}
