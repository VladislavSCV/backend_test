package handlers

import (
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Login - вход в систему
// @Summary Аутентификация пользователя
// @Description Аутентифицирует пользователя и возвращает JWT-токен
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param input body models.User true "Данные для входа"
// @Success 200 {object} map[string]string "Успешный вход, возвращает токен"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 401 {object} map[string]string "Неверные учетные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Получаем пользователя из базы данных
	dbUser, err := core.AuthenticateUser(db, user.Login, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Генерация токена
	token, err := utils.GenerateToken(dbUser.ID, dbUser.RoleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Registration - регистрация пользователя
// @Summary Регистрация нового пользователя
// @Description Создает нового пользователя в системе
// @Tags Аутентификация
// @Accept json
// @Produce json
// @Param input body models.User true "Данные для регистрации"
// @Success 200 {object} map[string]interface{} "Успешная регистрация, возвращает ID пользователя"
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/auth/registration [post]
func Registration(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Регистрация пользователя
	userID, err := core.RegisterUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user_id": userID})
}

// Verify - проверка токена
// @Summary Проверка JWT-токена
// @Description Проверяет валидность JWT-токена и возвращает информацию о пользователе
// @Tags Аутентификация
// @Produce json
// @Param Authorization header string true "JWT-токен"
// @Success 200 {object} map[string]interface{} "Успешная проверка, возвращает user_id и role_id"
// @Failure 401 {object} map[string]string "Неверный токен"
// @Router /api/auth/verify [post]
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

// GetCurrentUser - информация о текущем пользователе
// @Summary Получение информации о текущем пользователе
// @Description Возвращает информацию о текущем аутентифицированном пользователе
// @Tags Аутентификация
// @Produce json
// @Param Authorization header string true "JWT-токен"
// @Success 200 {object} map[string]interface{} "Информация о пользователе"
// @Failure 401 {object} map[string]string "Пользователь не аутентифицирован"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/auth/ [get]
func GetCurrentUser(c *gin.Context) {
	// Извлекаем userID из контекста
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	// Преобразуем userID в int
	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userID type"})
		return
	}

	// Получаем информацию о пользователе из базы данных
	user, err := core.GetCurrentUser(db, userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем информацию о пользователе
	c.JSON(http.StatusOK, gin.H{
		"id":          user.ID,
		"first_name":  user.FirstName,
		"middle_name": user.MiddleName,
		"last_name":   user.LastName,
		"role_id":     user.RoleID,
		"group_id":    user.GroupID,
		"login":       user.Login,
		"created_at":  user.CreatedAt,
		"updated_at":  user.UpdatedAt,
	})
}
