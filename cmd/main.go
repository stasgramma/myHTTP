package main

import (
	"log"

	"github.com/introxx/myhttp/config"
	"github.com/introxx/myhttp/internal/database"
	"github.com/introxx/myhttp/internal/models"
	"github.com/introxx/myhttp/internal/routes"

	_ "github.com/introxx/myhttp/docs" // для Swagger
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           MyApp API
// @version         1.0
// @description     Пример бэкенда на Go с Gin и Swagger
// @host            localhost:8080
// @BasePath        /
func main() {
	cfg := config.LoadConfig()
	database.Connect(cfg)

	// Создаём тестового пользователя
	createTestUser()

	r := routes.SetupRouter()

	// === 📘 Swagger UI ===
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Println("📘 Swagger доступен на http://localhost:8080/swagger/index.html")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

// Вне main() объявляем функцию
func createTestUser() {
	user := models.User{
		Name:  "Bob",
		Email: "bob@example.com",
		Role:  "user",
	}

	// Хешируем пароль
	if err := user.HashPassword("123456"); err != nil {
		log.Fatalf("Ошибка хеширования пароля: %v", err)
	}

	// Сохраняем в БД
	if err := database.DB.Create(&user).Error; err != nil {
		log.Fatalf("Ошибка создания пользователя: %v", err)
	}

	log.Println("✅ Пользователь создан:", user.Email)
	log.Println("Пароль в базе (должен быть хешем):", user.Password)

	// Проверка метода CheckPassword
	log.Println("Проверка пароля 123456:", user.CheckPassword("123456")) // true
	log.Println("Проверка пароля wrong:", user.CheckPassword("wrong"))   // false
}
