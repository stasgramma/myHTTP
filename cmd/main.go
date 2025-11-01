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
	"gorm.io/gorm"
)

// @title           MyApp API
// @version         1.0
// @description     Пример бэкенда на Go с Gin и Swagger
// @host            localhost:8080
// @BasePath        /
func main() {
	cfg := config.LoadConfig()
	database.Connect(cfg)

	// Создаём тестового пользователя, если его ещё нет
	createTestUser()

	r := routes.SetupRouter()

	// === 📘 Swagger UI ===
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Сервер запущен на http://localhost:8080")
	log.Println("Swagger доступен на http://localhost:8080/swagger/index.html")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

// createTestUser создаёт пользователя "Bob" один раз, если его нет
func createTestUser() {
	const email = "bob@example.com"
	var existing models.User

	// Проверяем, существует ли уже пользователь с таким email
	if err := database.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		log.Println("Пользователь уже существует, пропускаем создание.")
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Fatalf("Ошибка при проверке пользователя: %v", err)
	}

	// Создаём нового пользователя
	user := models.User{
		Name:  "Bob",
		Email: email,
		Role:  "user",
	}

	// Хешируем пароль перед сохранением
	if err := user.HashPassword("123456"); err != nil {
		log.Fatalf("Ошибка хеширования пароля: %v", err)
	}

	// Сохраняем пользователя в базе
	if err := database.DB.Create(&user).Error; err != nil {
		log.Fatalf("Ошибка создания пользователя: %v", err)
	}

	log.Println("Пользователь успешно создан:", user.Email)
	log.Println("Пароль сохранён в виде хеша:", user.Password)

	// Проверяем правильность пароля
	log.Println("Проверка правильного пароля (123456):", user.CheckPassword("123456"))
	log.Println("Проверка неверного пароля (wrong):", user.CheckPassword("wrong"))
}
