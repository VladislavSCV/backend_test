package handlers

import (
    "net/http"
    "github.com/VladislavSCV/internal/models"
    "github.com/gin-gonic/gin"
)

// GetGradesByStudentID - оценки студента
func GetGradesByStudentID(c *gin.Context) {
    studentID := c.Param("id")
    // Здесь будет логика получения оценок студента из базы данных
    grades := []models.Grade{} // Заглушка
    c.JSON(http.StatusOK, grades)
}

// GetGradesByGroupID - оценки группы
func GetGradesByGroupID(c *gin.Context) {
    groupID := c.Param("id")
    // Здесь будет логика получения оценок группы из базы данных
    grades := []models.Grade{} // Заглушка
    c.JSON(http.StatusOK, grades)
}

// CreateGrade - выставление оценки
func CreateGrade(c *gin.Context) {
    var grade models.Grade
    if err := c.ShouldBindJSON(&grade); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Здесь будет логика создания оценки в базе данных
    c.JSON(http.StatusOK, gin.H{"message": "Grade created"})
}

// UpdateGrade - обновление оценки
func UpdateGrade(c *gin.Context) {
    id := c.Param("id")
    var grade models.Grade
    if err := c.ShouldBindJSON(&grade); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Здесь будет логика обновления оценки в базе данных
    c.JSON(http.StatusOK, gin.H{"message": "Grade updated"})
}

// DeleteGrade - удаление оценки
func DeleteGrade(c *gin.Context) {
    id := c.Param("id")
    // Здесь будет логика удаления оценки из базы данных
    c.JSON(http.StatusOK, gin.H{"message": "Grade deleted"})
}