package handlers

import (
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetAttendanceByStudentID - посещаемость студента
// @Summary Получение посещаемости студента
// @Description Возвращает список отметок посещаемости для конкретного студента
// @Tags Посещаемость
// @Produce json
// @Param id path int true "ID студента"
// @Success 200 {array} models.Attendance "Список отметок посещаемости"
// @Failure 400 {object} map[string]string "Неверный ID студента"
// @Failure 404 {object} map[string]string "Посещаемость не найдена"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/attendance/student/{id} [get]
func GetAttendanceByStudentID(c *gin.Context) {
	studentID := c.Param("id")

	// Преобразуем studentID в int
	studentIDInt, err := strconv.Atoi(studentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
		return
	}

	attendances, err := core.GetAttendanceByStudentID(db, studentIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "attendance not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, attendances)
}

// GetAttendanceByGroupID - посещаемость группы
// @Summary Получение посещаемости группы
// @Description Возвращает список отметок посещаемости для всех студентов в группе
// @Tags Посещаемость
// @Produce json
// @Param id path int true "ID группы"
// @Success 200 {array} models.Attendance "Список отметок посещаемости"
// @Failure 400 {object} map[string]string "Неверный ID группы"
// @Failure 404 {object} map[string]string "Посещаемость не найдена"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/attendance/group/{id} [get]
func GetAttendanceByGroupID(c *gin.Context) {
	groupID := c.Param("id")

	// Преобразуем groupID в int
	groupIDInt, err := strconv.Atoi(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	attendances, err := core.GetAttendanceByGroupID(db, groupIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "attendance not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, attendances)
}

// CreateAttendance - отметка посещаемости
// @Summary Создание отметки посещаемости
// @Description Создает новую отметку посещаемости для студента
// @Tags Посещаемость
// @Accept json
// @Produce json
// @Param input body models.Attendance true "Данные для создания отметки посещаемости"
// @Success 200 {object} map[string]interface{} "Успешное создание отметки"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/attendance/ [post]
func CreateAttendance(c *gin.Context) {
	var attendance models.Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendanceID, err := core.CreateAttendance(db, attendance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance created successfully", "attendance_id": attendanceID})
}

// UpdateAttendance - обновление посещаемости
// @Summary Обновление отметки посещаемости
// @Description Обновляет существующую отметку посещаемости
// @Tags Посещаемость
// @Accept json
// @Produce json
// @Param id path int true "ID отметки посещаемости"
// @Param input body models.Attendance true "Данные для обновления отметки посещаемости"
// @Success 200 {object} map[string]string "Успешное обновление отметки"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/attendance/{id} [put]
func UpdateAttendance(c *gin.Context) {
	id := c.Param("id")

	// Преобразуем id в int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var attendance models.Attendance
	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем ID отметки посещаемости из параметра запроса
	attendance.ID = idInt

	if err := core.UpdateAttendance(db, attendance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance updated successfully"})
}
