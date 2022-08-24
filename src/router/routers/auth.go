package routers

import (
	authController "gotaskapp/src/controllers/auth"
	"gotaskapp/src/middlewares"

	"github.com/gin-gonic/gin"
)

// Routers of authentication
func Auth(app *gin.Engine) {
	authGroup := app.Group("/auth")
	{
		authGroup.POST("/sign/up", authController.SignUp)
		authGroup.POST("/sign/in", authController.SignIn)
		authGroup.GET("/email/verify/:token", authController.EmailVerify)
		authGroup.POST("/password/reset", authController.PasswordReset)
		authGroup.POST("/request/password/reset", authController.RequestPasswordReset)
		authGroup.POST("/{:id}/token/renew", middlewares.Authenticate(), authController.TokenRenew)
	}
}
