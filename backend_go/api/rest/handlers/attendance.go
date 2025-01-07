package handlers

import (
	"database/sql"
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetAttendanceByStudentID godoc
// @Summary Получить посещаемость студента по его ID
// @Description Возвращает список посещаемости для конкретного студента
// @Tags Attendance
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID студента"  example(1)
// @Success 200 {array} models.Attendance "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 404 {object} ErrorResponse "Посещаемость не найдена"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/attendance/student/{id} [get]
func GetAttendanceByStudentID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.Param("id")

		// Преобразуем studentID в int
		studentIDInt, err := strconv.Atoi(studentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid student ID"})
			return
		}

		attendances, err := core.GetAttendanceByStudentID(db, studentIDInt)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "attendance not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, attendances)
	}
}

// GetAttendanceByGroupID godoc
// @Summary Получить посещаемость группы по её ID
// @Description Возвращает список посещаемости для конкретной группы
// @Tags Attendance
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID группы"  example(1)
// @Success 200 {array} models.Attendance "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 404 {object} ErrorResponse "Посещаемость не найдена"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/attendance/group/{id} [get]
func GetAttendanceByGroupID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")

		// Преобразуем groupID в int
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid group ID"})
			return
		}

		attendances, err := core.GetAttendanceByGroupID(db, groupIDInt)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "attendance not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, attendances)
	}
}

// CreateAttendance godoc
// @Summary Создать отметку посещаемости
// @Description Создаёт новую запись о посещаемости
// @Tags Attendance
// @Accept  json
// @Produce  json
// @Param   attendance  body  models.Attendance  true  "Данные о посещаемости"  example({"student_id": 1, "subject_id": 1, "date": "2023-10-01T00:00:00Z", "status": "present"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/attendance [post]
func CreateAttendance(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var attendance models.Attendance
		if err := c.ShouldBindJSON(&attendance); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		attendanceID, err := core.CreateAttendance(db, attendance)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{
			Message: "Attendance created successfully",
			Data:    gin.H{"attendance_id": attendanceID},
		})
	}
}

// UpdateAttendance godoc
// @Summary Обновить отметку посещаемости
// @Description Обновляет существующую запись о посещаемости
// @Tags Attendance
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID посещаемости"  example(1)
// @Param   attendance  body  models.Attendance  true  "Обновлённые данные о посещаемости"  example({"student_id": 1, "subject_id": 1, "date": "2023-10-01T00:00:00Z", "status": "absent"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/attendance/{id} [put]
func UpdateAttendance(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Преобразуем id в int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid ID"})
			return
		}

		var attendance models.Attendance
		if err := c.ShouldBindJSON(&attendance); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		// Устанавливаем ID отметки посещаемости из параметра запроса
		attendance.ID = idInt

		if err := core.UpdateAttendance(db, attendance); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Attendance updated successfully"})
	}
}
