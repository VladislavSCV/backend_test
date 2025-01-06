package handlers

import (
    "net/http"
    "github.com/VladislavSCV/internal/models"
    "github.com/gin-gonic/gin"
)

// GetGroups - список всех групп
func GetGroups(c *gin.Context) {
    // Здесь будет логика получения групп из базы данных
    groups := []models.Group{} // Заглушка
    c.JSON(http.StatusOK, groups)
}

// GetGroupByID - информация о группе
func GetGroupByID(c *gin.Context) {
    id := c.Param("id")
    // Здесь будет логика получения группы по ID из базы данных
    group := models.Group{} // Заглушка
    c.JSON(http.StatusOK, group)
}

// CreateGroup - создание группы
func CreateGroup(c *gin.Context) {
    var group models.Group
    if err := c.ShouldBindJSON(&group); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Здесь будет логика создания группы в базе данных
    c.JSON(http.StatusOK, gin.H{"message": "Group created"})
}

// UpdateGroup - обновление группы
func UpdateGroup(c *gin.Context) {
    id := c.Param("id")
    var group models.Group
    if err := c.ShouldBindJSON(&group); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Здесь будет логика обновления группы в базе данных
    c.JSON(http.StatusOK, gin.H{"message": "Group updated"})
}

// DeleteGroup - удаление группы
func DeleteGroup(c *gin.Context) {
    id := c.Param("id")
    // Здесь будет логика удаления группы из базы данных
    c.JSON(http.StatusOK, gin.H{"message": "Group deleted"})
}