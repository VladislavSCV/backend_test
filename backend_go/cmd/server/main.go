package main

import (
    "github.com/gin-gonic/gin"
    "github.com/VladislavSCV/api/rest/routes"
    "github.com/VladislavSCV/internal/database"
)

func main() {
    r := gin.Default()

    // Подключение к базе данных
    database.ConnectToDB()

    routes.SetupUserRoutes(r)
    routes.SetupAuthRoutes(r)
    routes.SetupGroupRoutes(r)
    routes.SetupScheduleRoutes(r)
    routes.SetupGradeRoutes(r)
    routes.SetupAttendanceRoutes(r)

    // Запуск сервера
    r.Run(":8080")
}