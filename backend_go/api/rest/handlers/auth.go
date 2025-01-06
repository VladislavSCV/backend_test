package handlers

import (
    "net/http"
    "time"
    "github.com/VladislavSCV/internal/models"
    "github.com/VladislavSCV/internal/utils"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Здесь должен быть код для поиска пользователя в базе данных
    // Например: db.Where("login = ?", user.Login).First(&user)

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, err := utils.GenerateToken(user.ID, user.RoleID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Registration(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }

    user.Password = string(hashedPassword)
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    // Здесь должен быть код для сохранения пользователя в базе данных
    // Например: db.Create(&user)

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Verify(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
        return
    }

    claims, err := utils.ValidateToken(token)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user_id": claims.UserID, "role_id": claims.RoleID})
}

func GetCurrentUser(c *gin.Context) {
    userID := c.MustGet("userID").(int)
    roleID := c.MustGet("roleID").(int)

    // Здесь должен быть код для получения информации о пользователе из базы данных
    // Например: db.First(&user, userID)

    c.JSON(http.StatusOK, gin.H{"user_id": userID, "role_id": roleID})
}