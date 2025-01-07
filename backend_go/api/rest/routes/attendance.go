package routes

import (
	"database/sql"
	"github.com/VladislavSCV/api/middleware"
	"github.com/VladislavSCV/api/rest/handlers"
	"github.com/gin-gonic/gin"
)

func SetupAttendanceRoutes(router *gin.Engine, db *sql.DB) {
	attendanceGroup := router.Group("/api/attendance")
	{
		attendanceGroup.GET("/student/:id", handlers.GetAttendanceByStudentID(db))
		attendanceGroup.GET("/group/:id", handlers.GetAttendanceByGroupID(db))
		attendanceGroup.POST("/", middleware.AuthMiddleware(db), handlers.CreateAttendance(db))
		attendanceGroup.PUT("/:id", middleware.AuthMiddleware(db), handlers.UpdateAttendance(db))
	}
}
