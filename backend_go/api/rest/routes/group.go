package routes

import (
	"database/sql"
	"github.com/VladislavSCV/api/middleware"
	"github.com/VladislavSCV/api/rest/handlers"
	"github.com/gin-gonic/gin"
)

func SetupGroupRoutes(router *gin.Engine, db *sql.DB) {
	groupGroup := router.Group("/api/group")
	{
		groupGroup.GET("/", handlers.GetGroups(db))
		groupGroup.GET("/:id", handlers.GetGroupByID(db)) // получение информации о группе (студенты)
		groupGroup.POST("/", middleware.AuthMiddleware(db), handlers.CreateGroup(db))
		groupGroup.PUT("/:id", middleware.AuthMiddleware(db), handlers.UpdateGroup(db))
		groupGroup.DELETE("/:id", middleware.AuthMiddleware(db), handlers.DeleteGroup(db))
	}
}
