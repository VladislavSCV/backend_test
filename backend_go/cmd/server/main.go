package main

import (
	"github.com/VladislavSCV/api/rest/routes"
	_ "github.com/VladislavSCV/docs" // Импортируйте сгенерированную документацию
	"github.com/VladislavSCV/internal/database"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Маршрут для Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Подключение к базе данных
	database.ConnectDB()

	routes.SetupUserRoutes(r)
	routes.SetupAuthRoutes(r)
	routes.SetupGroupRoutes(r)
	routes.SetupScheduleRoutes(r)
	routes.SetupGradeRoutes(r)
	routes.SetupAttendanceRoutes(r)

	// Запуск сервера
	r.Run(":8080")
}
