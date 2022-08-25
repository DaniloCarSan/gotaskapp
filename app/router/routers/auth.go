package routers

import (
	authController "gotaskapp/app/controllers/auth"

	"github.com/gin-gonic/gin"
)

// Routers of authentication
func Auth(app *gin.Engine) {
	authGroup := app.Group("/auth")
	{
		authGroup.POST("/sign/up", authController.SignUp)
		authGroup.POST("/sign/in", authController.SignIn)
		authGroup.GET("/email/verify/:token", authController.EmailVerify)
		authGroup.POST("/password/reset/:token", authController.PasswordReset)
		authGroup.POST("/request/password/reset", authController.RequestPasswordReset)
	}
}
