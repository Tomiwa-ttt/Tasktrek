package routes

import (
    "github.com/gin-gonic/gin"
    "tasktrek/controllers"
    "tasktrek/middleware" // ✅ added
)

func UserRoutes(router *gin.Engine) {
    router.POST("/signup", controllers.SignUp)
    router.POST("/login", controllers.Login)

    // ✅ Protected routes
    protected := router.Group("/api")
    protected.Use(middleware.JWTAuthMiddleware())
    {
        protected.GET("/profile", controllers.Profile)
    }
}

