package handlers

import (
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetGradesByStudentID - оценки студента
// @Summary Получение оценок студента
// @Description Возвращает список оценок для конкретного студента
// @Tags Оценки
// @Produce json
// @Param id path int true "ID студента"
// @Success 200 {array} models.Grade "Список оценок"
// @Failure 400 {object} map[string]string "Неверный ID студента"
// @Failure 404 {object} map[string]string "Оценки не найдены"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/grades/student/{id} [get]
func GetGradesByStudentID(c *gin.Context) {
	studentID := c.Param("id")

	// Преобразуем studentID в int
	studentIDInt, err := strconv.Atoi(studentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student ID"})
		return
	}

	grades, err := core.GetGradesByStudentID(db, studentIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "grades not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, grades)
}

// GetGradesByGroupID - оценки группы
// @Summary Получение оценок группы
// @Description Возвращает список оценок для всех студентов в группе
// @Tags Оценки
// @Produce json
// @Param id path int true "ID группы"
// @Success 200 {array} models.Grade "Список оценок"
// @Failure 400 {object} map[string]string "Неверный ID группы"
// @Failure 404 {object} map[string]string "Оценки не найдены"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/grades/group/{id} [get]
func GetGradesByGroupID(c *gin.Context) {
	groupID := c.Param("id")

	// Преобразуем groupID в int
	groupIDInt, err := strconv.Atoi(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	grades, err := core.GetGradesByGroupID(db, groupIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "grades not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, grades)
}

// CreateGrade - выставление оценки
// @Summary Создание оценки
// @Description Создает новую оценку для студента
// @Tags Оценки
// @Accept json
// @Produce json
// @Param input body models.Grade true "Данные для создания оценки"
// @Success 200 {object} map[string]interface{} "Успешное создание оценки, возвращает ID оценки"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/grades/ [post]
func CreateGrade(c *gin.Context) {
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gradeID, err := core.CreateGrade(db, grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade created successfully", "grade_id": gradeID})
}

// UpdateGrade - обновление оценки
// @Summary Обновление оценки
// @Description Обновляет существующую оценку
// @Tags Оценки
// @Accept json
// @Produce json
// @Param id path int true "ID оценки"
// @Param input body models.Grade true "Данные для обновления оценки"
// @Success 200 {object} map[string]string "Успешное обновление оценки"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/grades/{id} [put]
func UpdateGrade(c *gin.Context) {
	id := c.Param("id")

	// Преобразуем id в int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем ID оценки из параметра запроса
	grade.ID = idInt

	if err := core.UpdateGrade(db, grade); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade updated successfully"})
}

// DeleteGrade - удаление оценки
// @Summary Удаление оценки
// @Description Удаляет существующую оценку
// @Tags Оценки
// @Produce json
// @Param id path int true "ID оценки"
// @Success 200 {object} map[string]string "Успешное удаление оценки"
// @Failure 400 {object} map[string]string "Неверный ID оценки"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/grades/{id} [delete]
func DeleteGrade(c *gin.Context) {
	id := c.Param("id")

	// Преобразуем id в int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := core.DeleteGrade(db, idInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade deleted successfully"})
}
