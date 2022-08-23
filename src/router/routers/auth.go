package routers

import (
	"github.com/gin-gonic/gin"
)

// Routers of authentication
func Auth(app *gin.Engine) {
	auth := app.Group("/auth")
	{
		auth.POST("/sign/in", func(ctx *gin.Context) {})
		auth.POST("/sign/up", func(ctx *gin.Context) {})
		auth.POST("/password/reset", func(ctx *gin.Context) {})
		auth.POST("/token/renew", func(ctx *gin.Context) {})
	}
}
