package handlers

import (
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// GetGroups - список всех групп
// @Summary Получение списка всех групп
// @Description Возвращает список всех групп
// @Tags Группы
// @Produce json
// @Success 200 {array} models.Group "Список групп"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/group/ [get]
func GetGroups(c *gin.Context) {
	groups, err := core.GetAllGroups(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// GetGroupByID - информация о группе
// @Summary Получение информации о группе
// @Description Возвращает информацию о конкретной группе по её ID
// @Tags Группы
// @Produce json
// @Param id path int true "ID группы"
// @Success 200 {object} models.Group "Информация о группе"
// @Failure 400 {object} map[string]string "Неверный ID группы"
// @Failure 404 {object} map[string]string "Группа не найдена"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/group/{id} [get]
func GetGroupByID(c *gin.Context) {
	groupID := c.Param("id")

	// Преобразуем groupID в int
	groupIDInt, err := strconv.Atoi(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	group, err := core.GetGroupByID(db, groupIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, group)
}

// CreateGroup - создание группы
// @Summary Создание новой группы
// @Description Создает новую группу
// @Tags Группы
// @Accept json
// @Produce json
// @Param input body models.Group true "Данные для создания группы"
// @Success 200 {object} map[string]interface{} "Успешное создание группы, возвращает ID группы"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/group/ [post]
func CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID, err := core.CreateGroup(db, group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group created successfully", "group_id": groupID})
}

// UpdateGroup - обновление группы
// @Summary Обновление информации о группе
// @Description Обновляет информацию о существующей группе
// @Tags Группы
// @Accept json
// @Produce json
// @Param id path int true "ID группы"
// @Param input body models.Group true "Данные для обновления группы"
// @Success 200 {object} map[string]string "Успешное обновление группы"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/group/{id} [put]
func UpdateGroup(c *gin.Context) {
	groupID := c.Param("id")

	// Преобразуем groupID в int
	groupIDInt, err := strconv.Atoi(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем ID группы из параметра запроса
	group.ID = groupIDInt

	if err := core.UpdateGroup(db, group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group updated successfully"})
}

// DeleteGroup - удаление группы
// @Summary Удаление группы
// @Description Удаляет существующую группу
// @Tags Группы
// @Produce json
// @Param id path int true "ID группы"
// @Success 200 {object} map[string]string "Успешное удаление группы"
// @Failure 400 {object} map[string]string "Неверный ID группы"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/group/{id} [delete]
func DeleteGroup(c *gin.Context) {
	groupID := c.Param("id")

	// Преобразуем groupID в int
	groupIDInt, err := strconv.Atoi(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	if err := core.DeleteGroup(db, groupIDInt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
}
