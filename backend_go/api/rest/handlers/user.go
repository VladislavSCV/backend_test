package handlers

import (
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetUsers - список всех пользователей
// @Summary Получение списка всех пользователей
// @Description Возвращает список всех пользователей в системе
// @Tags Пользователи
// @Produce json
// @Success 200 {array} models.User "Список пользователей"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/user/ [get]
func GetUsers(c *gin.Context) {
	users, err := core.GetAllUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetStudents - список студентов
// @Summary Получение списка студентов
// @Description Возвращает список всех студентов в системе
// @Tags Пользователи
// @Produce json
// @Success 200 {array} models.User "Список студентов"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/user/students [get]
func GetStudents(c *gin.Context) {
	students, err := core.GetStudents(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// GetTeachers - список преподавателей
// @Summary Получение списка преподавателей
// @Description Возвращает список всех преподавателей в системе
// @Tags Пользователи
// @Produce json
// @Success 200 {array} models.User "Список преподавателей"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/user/teachers [get]
func GetTeachers(c *gin.Context) {
	teachers, err := core.GetTeachers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teachers)
}

// GetUserByID - информация о пользователе
// @Summary Получение информации о пользователе
// @Description Возвращает информацию о конкретном пользователе по его ID
// @Tags Пользователи
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} models.User "Информация о пользователе"
// @Failure 400 {object} map[string]string "Неверный ID пользователя"
// @Failure 404 {object} map[string]string "Пользователь не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/user/{id} [get]
func GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	// Преобразуем userID в int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := core.GetUserByID(db, userIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser - обновление пользователя
// @Summary Обновление информации о пользователе
// @Description Обновляет информацию о пользователе по его ID
// @Tags Пользователи
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param input body models.User true "Данные для обновления пользователя"
// @Success 200 {object} map[string]string "Успешное обновление пользователя"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/user/{id} [put]
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	// Преобразуем userID в int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем ID пользователя из параметра запроса
	user.ID = userIDInt

	if err := core.UpdateUser(db, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser - удаление пользователя
// @Summary Удаление пользователя
// @Description Удаляет пользователя по его ID
// @Tags Пользователи
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} map[string]string "Успешное удаление пользователя"
// @Failure 400 {object} map[string]string "Неверный ID пользователя"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	// Преобразуем userID в int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := core.DeleteUser(db, userIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
