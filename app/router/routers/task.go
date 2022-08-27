package routers

import (
	taskController "gotaskapp/app/controllers/task"
	"gotaskapp/app/middlewares"

	"github.com/gin-gonic/gin"
)

// Routers of task
func Task(app *gin.Engine) {
	taskGroup := app.Group("/tasks", middlewares.Authenticate())
	{
		taskGroup.POST("", taskController.Create)
		taskGroup.GET("/status/:id", taskController.All)
		taskGroup.GET("/:id", taskController.Select)
		taskGroup.PUT("", taskController.Update)
		taskGroup.DELETE("/:id", taskController.Delete)
	}
}
