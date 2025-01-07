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

// GetSchedules godoc
// @Summary Получить общее расписание
// @Description Возвращает список всех занятий
// @Tags Schedules
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Schedule "Успешный ответ"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/schedule [get]
func GetSchedules(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		schedules, err := core.GetAllSchedules(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, schedules)
	}
}

// GetScheduleByID godoc
// @Summary Получить расписание по ID группы/преподавателя
// @Description Возвращает расписание для конкретной группы или преподавателя
// @Tags Schedules
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID группы или преподавателя"  example(1)
// @Success 200 {array} models.Schedule "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 404 {object} ErrorResponse "Расписание не найдено"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/schedule/{id} [get]
func GetScheduleByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Преобразуем id в int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid ID"})
			return
		}

		schedules, err := core.GetScheduleByID(db, idInt)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "schedule not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, schedules)
	}
}

// CreateSchedule godoc
// @Summary Создать занятие
// @Description Создаёт новое занятие в расписании
// @Tags Schedules
// @Accept  json
// @Produce  json
// @Param   schedule  body  models.Schedule  true  "Данные о занятии"  example({"group_id": 1, "teacher_id": 1, "subject_id": 1, "start_time": "2023-10-01T09:00:00Z", "end_time": "2023-10-01T10:30:00Z"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/schedule [post]
func CreateSchedule(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var schedule models.Schedule
		if err := c.ShouldBindJSON(&schedule); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		scheduleID, err := core.CreateSchedule(db, schedule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{
			Message: "Schedule created successfully",
			Data:    gin.H{"schedule_id": scheduleID},
		})
	}
}

// UpdateSchedule godoc
// @Summary Обновить занятие
// @Description Обновляет информацию о существующем занятии
// @Tags Schedules
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID занятия"  example(1)
// @Param   updates  body  map[string]interface{}  true  "Обновлённые данные о занятии"  example({"start_time": "2023-10-01T10:00:00Z", "end_time": "2023-10-01T11:30:00Z"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/schedule/{id} [put]
func UpdateSchedule(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Преобразуем id в int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid ID"})
			return
		}

		var updates map[string]interface{}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		// Устанавливаем ID занятия из параметра запроса
		updates["id"] = idInt

		if err := core.UpdateSchedule(db, updates); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Schedule updated successfully"})
	}
}

// DeleteSchedule godoc
// @Summary Удалить занятие
// @Description Удаляет занятие по его ID
// @Tags Schedules
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID занятия"  example(1)
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/schedule/{id} [delete]
func DeleteSchedule(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Преобразуем id в int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid ID"})
			return
		}

		if err := core.DeleteSchedule(db, idInt); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Schedule deleted successfully"})
	}
}
