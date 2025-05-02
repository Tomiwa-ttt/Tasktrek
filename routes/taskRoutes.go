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
	taskGroup.PATCH("/:id", controllers.UpdateTask) // Add the new route for updating tasks
}
