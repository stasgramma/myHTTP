package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/introxx/myhttp/internal/database"
	"github.com/introxx/myhttp/internal/models"
)

// PingHandler godoc
// @Summary Проверка доступности сервера
// @Description Возвращает pong, если сервер работает
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// GetUsers godoc
// @Summary Получить список пользователей
// @Description Возвращает всех пользователей из базы данных
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения пользователей"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Создать нового пользователя
// @Description Создаёт пользователя и сохраняет его в базе данных
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения пользователя"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
