package main

import (
	"log"

	"github.com/introxx/myhttp/config"
	"github.com/introxx/myhttp/internal/database"
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

	r := routes.SetupRouter()

	// === 📘 Swagger UI ===
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("🚀 Сервер запущен на http://localhost:8080")
	log.Println("📘 Swagger доступен на http://localhost:8080/swagger/index.html")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
