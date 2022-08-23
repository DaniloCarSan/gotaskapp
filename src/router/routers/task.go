package routers

import "github.com/gin-gonic/gin"

// Routers of task
func Task(app *gin.Engine) {
	tasks := app.Group("/tasks")
	{
		tasks.POST("", func(ctx *gin.Context) {})
		tasks.GET("", func(ctx *gin.Context) {})
		tasks.GET("/:id", func(ctx *gin.Context) {})
		tasks.PUT("", func(ctx *gin.Context) {})
		tasks.DELETE("/:id", func(ctx *gin.Context) {})
		tasks.PATCH("/toggle/done/{id:[0-9]+}", func(ctx *gin.Context) {})
	}
}
