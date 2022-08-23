package routers

import (
	"gotaskapp/src/middlewares"

	"github.com/gin-gonic/gin"
)

// Routers of user
func User(app *gin.Engine) {
	users := app.Group("/users", middlewares.Authenticate())
	{
		users.GET("", func(ctx *gin.Context) {})
		users.PATCH("", func(ctx *gin.Context) {})
	}
}
