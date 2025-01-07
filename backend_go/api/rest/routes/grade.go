package routes

import (
	"database/sql"
	"github.com/VladislavSCV/api/middleware"
	"github.com/VladislavSCV/api/rest/handlers"
	"github.com/gin-gonic/gin"
)

func SetupGradeRoutes(router *gin.Engine, db *sql.DB) {
	gradeGroup := router.Group("/api/grades")
	{
		gradeGroup.GET("/student/:id", handlers.GetGradesByStudentID(db))
		gradeGroup.GET("/group/:id", handlers.GetGradesByGroupID(db))
		gradeGroup.POST("/", middleware.AuthMiddleware(db), handlers.CreateGrade(db))
		gradeGroup.PUT("/:id", middleware.AuthMiddleware(db), handlers.UpdateGrade(db))
		gradeGroup.DELETE("/:id", middleware.AuthMiddleware(db), handlers.DeleteGrade(db))
	}
}
