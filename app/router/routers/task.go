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
		taskGroup.PUT("/:id", taskController.Update)
		taskGroup.DELETE("/:id", taskController.Delete)
		taskGroup.PATCH("/:task_id/status/:status_id", taskController.ChangeStatus)
	}
}
