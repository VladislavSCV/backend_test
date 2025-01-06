package routes

import (
    "github.com/VladislavSCV/api/rest/handlers"
    "github.com/VladislavSCV/api/middleware"
    "github.com/gin-gonic/gin"
)

func SetupGroupRoutes(router *gin.Engine) {
    groupGroup := router.Group("/api/group")
    {
        groupGroup.GET("/", handlers.GetGroups)
        groupGroup.GET("/:id", handlers.GetGroupByID)
        groupGroup.POST("/", middleware.AuthMiddleware(), handlers.CreateGroup)
        groupGroup.PUT("/:id", middleware.AuthMiddleware(), handlers.UpdateGroup)
        groupGroup.DELETE("/:id", middleware.AuthMiddleware(), handlers.DeleteGroup)
    }
}