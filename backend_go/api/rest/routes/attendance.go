package routes

import (
    "github.com/VladislavSCV/api/rest/handlers"
    "github.com/VladislavSCV/api/middleware"
    "github.com/gin-gonic/gin"
)

func SetupAttendanceRoutes(router *gin.Engine) {
    attendanceGroup := router.Group("/api/attendance")
    {
        attendanceGroup.GET("/student/:id", handlers.GetAttendanceByStudentID)
        attendanceGroup.GET("/group/:id", handlers.GetAttendanceByGroupID)
        attendanceGroup.POST("/", middleware.AuthMiddleware(), handlers.CreateAttendance)
        attendanceGroup.PUT("/:id", middleware.AuthMiddleware(), handlers.UpdateAttendance)
    }
}