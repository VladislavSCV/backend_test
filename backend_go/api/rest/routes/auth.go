package routes

import (
    "github.com/VladislavSCV/api/rest/handlers"
    "github.com/VladislavSCV/api/middleware"
    "github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
    // Аутентификация
    auth := router.Group("/api/auth")
    {
        auth.POST("/login", handlers.Login)
        auth.POST("/registration", handlers.Registration)
        auth.POST("/verify", handlers.Verify)
        auth.GET("/", middleware.AuthMiddleware(), handlers.GetCurrentUser)
    }

}