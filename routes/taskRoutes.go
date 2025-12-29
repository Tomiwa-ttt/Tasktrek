package routes

import (
	"tasktrek/controllers"
	"tasktrek/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/api/tasks")
	taskGroup.Use(middleware.AuthMiddleware()) // Protect routes with JWT

	taskGroup.POST("/", controllers.CreateTask)
	taskGroup.GET("/", controllers.GetTasks)
	taskGroup.PATCH("/:id", controllers.UpdateTask)
	taskGroup.DELETE("/:id", controllers.DeleteTask)
}
