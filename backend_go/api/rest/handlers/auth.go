package handlers

import (
	"database/sql"
	"github.com/VladislavSCV/internal/core"
	"github.com/VladislavSCV/internal/models"
	"github.com/VladislavSCV/internal/utils"
	"github.com/VladislavSCV/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login godoc
// @Summary Вход в систему
// @Description Аутентификация пользователя и получение токена
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "Данные для входа"  example({"login": "user123", "password": "password123"})
// @Success 200 {object} LoginResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 401 {object} ErrorResponse "Неверные учетные данные"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/auth/login [post]
func Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		dbUser, err := core.AuthenticateUser(db, user.Login, user.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials"})
			return
		}

		token, err := utils.GenerateToken(dbUser.ID, dbUser.RoleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, LoginResponse{Token: token})
	}
}

// Registration godoc
// @Summary Регистрация пользователя
// @Description Создание нового пользователя в системе
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "Данные для регистрации"  example({"login": "newuser", "password": "newpassword123", "first_name": "John", "last_name": "Doe", "role_id": 1})
// @Success 200 {object} SuccessResponse "Успешный ответ"
// @Failure 400 {object} ErrorResponse "Неверный запрос"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/auth/registration [post]
func Registration(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
			return
		}

		hashResult, err := pkg.CreateHashWithSalt(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Could not hash password"})
			return
		}

		user.Password = hashResult.Hash
		user.Salt = hashResult.Salt
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		userID, err := core.RegisterUser(db, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, SuccessResponse{
			Message: "User registered successfully",
			Data:    gin.H{"user_id": userID},
		})
	}
}

// Verify godoc
// @Summary Проверка токена
// @Description Проверка валидности токена и получение информации о пользователе
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Токен авторизации"  example("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
// @Success 200 {object} VerifyResponse "Успешный ответ"
// @Failure 401 {object} ErrorResponse "Неверный токен"
// @Router /api/auth/verify [post]
func Verify(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "No token provided"})
			return
		}

		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid token"})
			return
		}

		c.JSON(http.StatusOK, VerifyResponse{
			UserID: claims.UserID,
			RoleID: claims.RoleID,
		})
	}
}

// GetCurrentUser godoc
// @Summary Информация о текущем пользователе
// @Description Получение информации о текущем аутентифицированном пользователе
// @Tags User
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Токен авторизации"  example("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...")
// @Success 200 {object} UserResponse "Успешный ответ"
// @Failure 401 {object} ErrorResponse "Пользователь не аутентифицирован"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Router /api/auth [get]
func GetCurrentUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "userID not found in context"})
			return
		}

		userIDInt, ok := userID.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "invalid userID type"})
			return
		}

		user, err := core.GetCurrentUser(db, userIDInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, UserResponse{
			ID:         user.ID,
			FirstName:  user.FirstName,
			MiddleName: user.MiddleName,
			LastName:   user.LastName,
			RoleID:     user.RoleID,
			GroupID:    user.GroupID,
			Login:      user.Login,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		})
	}
}
