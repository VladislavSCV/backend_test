package routes

import (
    "github.com/VladislavSCV/api/rest/handlers"
    "github.com/VladislavSCV/api/middleware"
    "github.com/gin-gonic/gin"
)

func SetupGradeRoutes(router *gin.Engine) {
    gradeGroup := router.Group("/api/grades")
    {
        gradeGroup.GET("/student/:id", handlers.GetGradesByStudentID)
        gradeGroup.GET("/group/:id", handlers.GetGradesByGroupID)
        gradeGroup.POST("/", middleware.AuthMiddleware(), handlers.CreateGrade)
        gradeGroup.PUT("/:id", middleware.AuthMiddleware(), handlers.UpdateGrade)
        gradeGroup.DELETE("/:id", middleware.AuthMiddleware(), handlers.DeleteGrade)
    }
}