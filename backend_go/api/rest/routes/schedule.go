package routes

import (
	"database/sql"
	"github.com/VladislavSCV/api/middleware"
	"github.com/VladislavSCV/api/rest/handlers"
	"github.com/gin-gonic/gin"
)

func SetupScheduleRoutes(router *gin.Engine, db *sql.DB) {
	scheduleGroup := router.Group("/api/schedule")
	{
		scheduleGroup.GET("/", handlers.GetSchedules(db))
		scheduleGroup.GET("/:id", handlers.GetScheduleByID(db))
		scheduleGroup.POST("/", middleware.AuthMiddleware(db), handlers.CreateSchedule(db))
		scheduleGroup.PUT("/:id", middleware.AuthMiddleware(db), handlers.UpdateSchedule(db))
		scheduleGroup.DELETE("/:id", middleware.AuthMiddleware(db), handlers.DeleteSchedule(db))
	}
}
