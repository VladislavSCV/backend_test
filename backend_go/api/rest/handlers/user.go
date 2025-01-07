package handlers

import (
	"database/sql"
	"github.com/VladislavSCV/internal/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetUsers godoc
// @Summary Получить список всех пользователей
// @Description Возвращает список всех пользователей
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User "Успешный ответ"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user [get]
func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := core.GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

// GetStudents godoc
// @Summary Получить список студентов
// @Description Возвращает список всех студентов
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User "Успешный ответ"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/students [get]
func GetStudents(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		students, err := core.GetStudents(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, students)
	}
}

// GetTeachers godoc
// @Summary Получить список преподавателей
// @Description Возвращает список всех преподавателей
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User "Успешный ответ"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/teachers [get]
func GetTeachers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		teachers, err := core.GetTeachers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, teachers)
	}
}

// GetUserByID godoc
// @Summary Получить информацию о пользователе по его ID
// @Description Возвращает информацию о конкретном пользователе
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID пользователя"  example(1)
// @Success 200 {object} models.User "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 404 {object} ErrorResponse "Пользователь не найден"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/{id} [get]
func GetUserByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		// Преобразуем userID в int
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user ID"})
			return
		}

		user, err := core.GetUserByID(db, userIDInt)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "user not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser godoc
// @Summary Обновить информацию о пользователе
// @Description Обновляет информацию о существующем пользователе
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID пользователя"  example(1)
// @Param   updates  body  map[string]interface{}  true  "Обновлённые данные о пользователе"  example({"first_name": "John", "last_name": "Doe"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/{id} [put]
func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		// Преобразуем userID в int
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user ID"})
			return
		}

		var updates map[string]interface{}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		// Устанавливаем ID пользователя из параметра запроса
		updates["id"] = userIDInt

		if err := core.UpdateUser(db, updates); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "User updated successfully"})
	}
}

// DeleteUser godoc
// @Summary Удалить пользователя
// @Description Удаляет пользователя по его ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID пользователя"  example(1)
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/user/{id} [delete]
func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		// Преобразуем userID в int
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid user ID"})
			return
		}

		if err := core.DeleteUser(db, userIDInt); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "User deleted successfully"})
	}
}
