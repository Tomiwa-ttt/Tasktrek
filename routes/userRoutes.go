package routes

import (
	"tasktrek/controllers"
	"tasktrek/middleware" // ✅ added

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	// ✅ Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", controllers.Profile)
	}
}
