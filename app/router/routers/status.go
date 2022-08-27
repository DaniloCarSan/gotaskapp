package routers

import (
	statusController "gotaskapp/app/controllers/status"
	"gotaskapp/app/middlewares"

	"github.com/gin-gonic/gin"
)

// Routers of status
func Status(app *gin.Engine) {
	statusGroup := app.Group("/status", middlewares.Authenticate())
	{
		statusGroup.POST("", statusController.Create)
		statusGroup.GET("", statusController.All)
	}
}
