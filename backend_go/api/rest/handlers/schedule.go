package handlers

import (
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetSchedules - общее расписание
// @Summary Получение общего расписания
// @Description Возвращает список всех занятий в расписании
// @Tags Расписание
// @Produce json
// @Success 200 {array} models.Schedule "Список занятий"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/schedule/ [get]
func GetSchedules(c *gin.Context) {
	schedules, err := core.GetAllSchedules(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetScheduleByID - расписание группы/преподавателя
// @Summary Получение расписания группы или преподавателя
// @Description Возвращает список занятий для конкретной группы или преподавателя
// @Tags Расписание
// @Produce json
// @Param id path int true "ID группы или преподавателя"
// @Success 200 {array} models.Schedule "Список занятий"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Расписание не найдено"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/schedule/{id} [get]
func GetScheduleByID(c *gin.Context) {
	id := c.Param("id")

	// Преобразуем id в int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	schedules, err := core.GetScheduleByID(db, idInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "schedule not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// CreateSchedule - создание занятия
// @Summary Создание нового занятия
// @Description Создает новое занятие в расписании
// @Tags Расписание
// @Accept json
// @Produce json
// @Param input body models.Schedule true "Данные для создания занятия"
// @Success 200 {object} map[string]interface{} "Успешное создание занятия, возвращает ID занятия"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/schedule/ [post]
func CreateSchedule(c *gin.Context) {
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheduleID, err := core.CreateSchedule(db, schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule created successfully", "schedule_id": scheduleID})
}

// UpdateSchedule - обновление занятия
// @Summary Обновление существующего занятия
// @Description Обновляет данные существующего занятия в расписании
// @Tags Расписание
// @Accept json
// @Produce json
// @Param id path int true "ID занятия"
// @Param input body models.Schedule true "Данные для обновления занятия"
// @Success 200 {object} map[string]string "Успешное обновление занятия"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/schedule/{id} [put]
func UpdateSchedule(c *gin.Context) {
	id := c.Param("id")

	// Преобразуем id в int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем ID занятия из параметра запроса
	schedule.ID = idInt

	if err := core.UpdateSchedule(db, schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
}

// DeleteSchedule - удаление занятия
// @Summary Удаление занятия
// @Description Удаляет занятие из расписания
// @Tags Расписание
// @Produce json
// @Param id path int true "ID занятия"
// @Success 200 {object} map[string]string "Успешное удаление занятия"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/schedule/{id} [delete]
func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")

	// Преобразуем id в int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := core.DeleteSchedule(db, idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule deleted successfully"})
}
