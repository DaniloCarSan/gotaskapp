package routers

import "github.com/gin-gonic/gin"

// Routers of user
func User(app *gin.Engine) {
	users := app.Group("/users")
	{
		users.GET("", func(ctx *gin.Context) {})
		users.PATCH("", func(ctx *gin.Context) {})
	}
}
