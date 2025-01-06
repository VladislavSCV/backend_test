package routes

import (
    "github.com/VladislavSCV/api/rest/handlers"
    "github.com/VladislavSCV/api/middleware"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
    userGroup := router.Group("/api/user")
    {
        userGroup.GET("/", handlers.GetUsers)
        userGroup.GET("/students", handlers.GetStudents)
        userGroup.GET("/teachers", handlers.GetTeachers)
        userGroup.GET("/:id", handlers.GetUserByID)
        userGroup.PUT("/:id", middleware.AuthMiddleware(), handlers.UpdateUser)
        userGroup.DELETE("/:id", middleware.AuthMiddleware(), handlers.DeleteUser)
    }
}