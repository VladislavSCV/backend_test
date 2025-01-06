package routes

import (
    "github.com/VladislavSCV/api/rest/handlers"
    "github.com/VladislavSCV/api/middleware"
    "github.com/gin-gonic/gin"
)

func SetupScheduleRoutes(router *gin.Engine) {
    scheduleGroup := router.Group("/api/schedule")
    {
        scheduleGroup.GET("/", handlers.GetSchedules)
        scheduleGroup.GET("/:id", handlers.GetScheduleByID)
        scheduleGroup.POST("/", middleware.AuthMiddleware(), handlers.CreateSchedule)
        scheduleGroup.PUT("/:id", middleware.AuthMiddleware(), handlers.UpdateSchedule)
        scheduleGroup.DELETE("/:id", middleware.AuthMiddleware(), handlers.DeleteSchedule)
    }
}