package routers

import (
	taskController "gotaskapp/src/controllers/task"
	"gotaskapp/src/middlewares"

	"github.com/gin-gonic/gin"
)

// Routers of task
func Task(app *gin.Engine) {
	taskGroup := app.Group("/tasks", middlewares.Authenticate())
	{
		taskGroup.POST("", taskController.Create)
		taskGroup.GET("", taskController.All)
		taskGroup.GET("/:id", taskController.Select)
		taskGroup.PUT("", taskController.Update)
		taskGroup.DELETE("/:id", taskController.Delete)
		taskGroup.PATCH("/toggle/done/{id}", taskController.Done)
	}
}
