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

// GetGroups godoc
// @Summary Получить список всех групп
// @Description Возвращает список всех групп
// @Tags Groups
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Group "Успешный ответ"  example([{"id": 1, "name": "Group A"}, {"id": 2, "name": "Group B"}])
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/group [get]
func GetGroups(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groups, err := core.GetAllGroups(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, groups)
	}
}

// GetGroupByID godoc
// @Summary Получить информацию о группе по её ID
// @Description Возвращает информацию о конкретной группе
// @Tags Groups
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID группы"  example(1)
// @Success 200 {object} models.Group "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 404 {object} ErrorResponse "Группа не найдена"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/group/{id} [get]
func GetGroupByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")

		// Преобразуем groupID в int
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid group ID"})
			return
		}

		group, err := core.GetGroupByID(db, groupIDInt)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.JSON(http.StatusNotFound, ErrorResponse{Error: "group not found"})
			} else {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, group)
	}
}

// CreateGroup godoc
// @Summary Создать группу
// @Description Создаёт новую группу
// @Tags Groups
// @Accept  json
// @Produce  json
// @Param   group  body  models.Group  true  "Данные о группе"  example({"name": "Group A"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/group [post]
func CreateGroup(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var group models.Group
		if err := c.ShouldBindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		groupID, err := core.CreateGroup(db, group)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{
			Message: "Group created successfully",
			Data:    gin.H{"group_id": groupID},
		})
	}
}

// UpdateGroup godoc
// @Summary Обновить информацию о группе
// @Description Обновляет информацию о существующей группе
// @Tags Groups
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID группы"  example(1)
// @Param   group  body  models.Group  true  "Обновлённые данные о группе"  example({"name": "Updated Group A"})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/group/{id} [put]
func UpdateGroup(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")

		// Преобразуем groupID в int
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid group ID"})
			return
		}

		var group models.Group
		if err := c.ShouldBindJSON(&group); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		// Устанавливаем ID группы из параметра запроса
		group.ID = groupIDInt

		if err := core.UpdateGroup(db, group); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Group updated successfully"})
	}
}

// DeleteGroup godoc
// @Summary Удалить группу
// @Description Удаляет группу по её ID
// @Tags Groups
// @Accept  json
// @Produce  json
// @Param   id  path  int  true  "ID группы"  example(1)
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/group/{id} [delete]
func DeleteGroup(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		groupID := c.Param("id")

		// Преобразуем groupID в int
		groupIDInt, err := strconv.Atoi(groupID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid group ID"})
			return
		}

		if err := core.DeleteGroup(db, groupIDInt); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{Message: "Group deleted successfully"})
	}
}
