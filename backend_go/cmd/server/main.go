package main

import (
	"database/sql"
	"fmt"
	"github.com/VladislavSCV/api/rest/routes"
	_ "github.com/VladislavSCV/docs" // Импортируйте сгенерированную документацию
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	// Подключение к базе данных
	connStr := "postgresql://test_db_mvaz_user:NgcEd82NG6iHSgqfwhkhPukcnsBHC0c4@dpg-cttsnalumphs73ei09c0-a.oregon-postgres.render.com/test_db_mvaz"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Проверка соединения с базой данных
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
	fmt.Println("Database connection:", db)

	// Инициализация Gin
	r := gin.Default()

	// Применение rate limiting ко всем API endpoints
	//apiGroup := r.Group("/api")
	//apiGroup.Use(middleware.RateLimiterMiddleware())

	// Передача db в маршруты
	routes.SetupAuthRoutes(r, db)
	routes.SetupUserRoutes(r, db)
	routes.SetupGroupRoutes(r, db)
	routes.SetupScheduleRoutes(r, db)
	routes.SetupGradeRoutes(r, db)
	routes.SetupAttendanceRoutes(r, db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Маршрут для Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
