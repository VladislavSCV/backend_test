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

// GetGradesByStudentID godoc
// @Summary Получить оценки студента по его ID
// @Description Возвращает список оценок для конкретного студента
// @Tags Grades
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID студента"  example(1)
// @Success 200 {array} models.Grade "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 404 {object} ErrorResponse "Оценки не найдены"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/grades/student/{id} [get]
func GetGradesByStudentID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.Param("id")

		// Преобразуем studentID в int
		studentIDInt, err := strconv.Atoi(studentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid student ID"})
			return
		}

		grades, err := core.GetGradesByStudentID(db, studentIDInt)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "grades not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, grades)
	}
}

// GetGradesByGroupID godoc
// @Summary Получить оценки группы по её ID
// @Description Возвращает список оценок для конкретной группы
// @Tags Grades
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID группы"  example(1)
// @Success 200 {array} models.Grade "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 404 {object} ErrorResponse "Оценки не найдены"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/grades/group/{id} [get]
func GetGradesByGroupID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")

		// Преобразуем groupID в int
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid group ID"})
			return
		}

		grades, err := core.GetGradesByGroupID(db, groupIDInt)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "grades not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, grades)
	}
}

// CreateGrade godoc
// @Summary Создать оценку
// @Description Создаёт новую запись об оценке
// @Tags Grades
// @Accept  json
// @Produce  json
// @Param   grade  body  models.Grade  true  "Данные об оценке"  example({"student_id": 1, "subject_id": 1, "grade": 5, "date": "2023-10-01T00:00:00Z"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/grades [post]
func CreateGrade(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var grade models.Grade
		if err := c.ShouldBindJSON(&grade); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		gradeID, err := core.CreateGrade(db, grade)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{
			Message: "Grade created successfully",
			Data:    gin.H{"grade_id": gradeID},
		})
	}
}

// UpdateGrade godoc
// @Summary Обновить оценку
// @Description Обновляет существующую запись об оценке
// @Tags Grades
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID оценки"  example(1)
// @Param   grade  body  models.Grade  true  "Обновлённые данные об оценке"  example({"student_id": 1, "subject_id": 1, "grade": 4, "date": "2023-10-01T00:00:00Z"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/grades/{id} [put]
func UpdateGrade(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Преобразуем id в int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid ID"})
			return
		}

		var grade models.Grade
		if err := c.ShouldBindJSON(&grade); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		// Устанавливаем ID оценки из параметра запроса
		grade.ID = idInt

		if err := core.UpdateGrade(db, grade); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Grade updated successfully"})
	}
}

// DeleteGrade godoc
// @Summary Удалить оценку
// @Description Удаляет запись об оценке
// @Tags Grades
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID оценки"  example(1)
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/grades/{id} [delete]
func DeleteGrade(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Преобразуем id в int
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid ID"})
			return
		}

		if err := core.DeleteGrade(db, idInt); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Grade deleted successfully"})
	}
}
